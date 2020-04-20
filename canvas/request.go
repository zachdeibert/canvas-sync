package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/zachdeibert/canvas-sync/task"
)

var (
	linkRe = regexp.MustCompile("<([^>]+)>;\\s*rel=\"([^\"]+)\"")
)

// Request sends a request to the API
func (c *Canvas) Request(endpoint string, params map[string]interface{}, progress *task.Progress, responseCtor func() interface{}, callback func(interface{}) error) error {
	sParams := make([]string, len(params))
	i := 0
	if params != nil {
		for k, v := range params {
			var err error
			if sParams[i], err = c.serializeParameter(k, v); err != nil {
				return err
			}
			i++
		}
	}
	url := fmt.Sprintf("https://%s.instructure.com/api/v1/%s?%s", c.subdomain, endpoint, strings.Join(sParams, "&"))
	progress.SetWork(1)
	first := true
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
		req.Header.Add("Accept", "application/json")
		c.onRequestStart()
		res, err := c.client.Do(req)
		c.onRequestFinish(res, err)
		if err != nil {
			return err
		}
		if res.StatusCode < 200 || res.StatusCode >= 300 {
			return fmt.Errorf("Invalid status code at URL %s: %s", url, res.Status)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		response := responseCtor()
		if err = json.Unmarshal(body, response); err != nil {
			return err
		}
		if err = callback(response); err != nil {
			return err
		}
		progress.Finish(1)
		links := linkRe.FindAllStringSubmatch(res.Header.Get("Link"), -1)
		if links == nil {
			return nil
		}
		found := false
		for _, link := range links {
			switch link[2] {
			case "next":
				url = link[1]
				found = true
				break
			case "last":
				if first {
					query := strings.Split(strings.Split(link[1], "?")[1], "&")
					for _, param := range query {
						parts := strings.Split(param, "=")
						if parts[0] == "page" {
							var numPages int
							fmt.Sscanf(parts[1], "%d", &numPages)
							progress.SetWork(numPages)
							first = false
							break
						}
					}
				}
				break
			}
		}
		if first && found {
			progress.AddWork(1)
		}
		if !found {
			return nil
		}
	}
}
