package canvas

import (
	"encoding/json"
	"errors"
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

// InvalidStatusCodeError is an error when the server responds with an invalid status code
type InvalidStatusCodeError struct {
	URL    string
	Status string
	Code   int
	Body   string
}

func (err InvalidStatusCodeError) Error() string {
	return fmt.Sprintf("Invalid status code at URL %s: %s", err.URL, err.Status)
}

// RequestRaw performs a raw HTTP request
func (c *Canvas) RequestRaw(url string, accept string, allowedRedirects int) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	if req.URL.Host == fmt.Sprintf("%s.instructure.com", c.subdomain) {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	req.Header.Add("Accept", accept)
	c.onRequestStart()
	res, err := c.client.Do(req)
	c.onRequestFinish(res, err)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, nil, InvalidStatusCodeError{
			URL:    url,
			Status: res.Status,
			Code:   res.StatusCode,
			Body:   string(body),
		}
	}
	if res.StatusCode >= 300 {
		loc := res.Header.Get("Location")
		if len(loc) > 0 {
			if allowedRedirects > 0 {
				body, res, err := c.RequestRaw(loc, accept, allowedRedirects-1)
				if err == nil {
					err = c.saveRequest(*req.URL, fmt.Sprintf("%s+%s", accept, RedirectType), []byte(loc), res)
				}
				return body, res, err
			}
			return nil, nil, errors.New("Too many redirects")
		}
		return nil, nil, errors.New("Redirect requested with no target location")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	err = c.saveRequest(*req.URL, accept, body, res)
	return body, res, err
}

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
	progress.AddWork(1)
	first := true
	for i = 1; ; i++ {
		body, res, err := c.RequestRaw(url, "application/json", 10)
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
							progress.AddWork(numPages - i)
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
