package canvas

import (
	"fmt"
	"net/http"
)

// Canvas represents the Canvas API
type Canvas struct {
	subdomain      string
	token          string
	parameterTypes []parameterType
	client         *http.Client
}

// CreateCanvas creates a new Canvas object
func CreateCanvas(subdomain, token string) (*Canvas, error) {
	c := &Canvas{
		subdomain:      subdomain,
		token:          token,
		parameterTypes: []parameterType{},
		client:         &http.Client{},
	}
	if err := c.registerDefaultParameterTypes(); err != nil {
		return nil, err
	}
	return c, nil
}

// GetSubdomain of the Canvas API
func (c *Canvas) GetSubdomain() string {
	return c.subdomain
}

// GetBaseURL of the Canvas API
func (c *Canvas) GetBaseURL() string {
	return fmt.Sprintf("https://%s.instructure.com/", c.subdomain)
}
