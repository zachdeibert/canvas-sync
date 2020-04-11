package canvas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	linkRe = regexp.MustCompile("<([^>]+)>;\\s*rel=\"([^\"]+)\"")
)

// Request sends a request to the API
func (c *Canvas) Request(endpoint string, params map[string]interface{}, responseCtor func() interface{}, callback func(interface{}) error) error {
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
	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
		req.Header.Add("Accept", "application/json")
		res, err := c.client.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode < 200 || res.StatusCode >= 300 {
			return fmt.Errorf("Invalid status code: %d %s", res.StatusCode, res.Status)
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
		links := linkRe.FindAllStringSubmatch(res.Header.Get("Link"), -1)
		if links == nil {
			return nil
		}
		found := false
		for _, link := range links {
			if link[2] == "next" {
				url = link[1]
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
}
