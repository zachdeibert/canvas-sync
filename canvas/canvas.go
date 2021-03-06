package canvas

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Canvas represents the Canvas API
type Canvas struct {
	subdomain       string
	token           string
	parameterTypes  []parameterType
	client          *http.Client
	lastQuota       float64
	lastQuotaTime   time.Time
	quotaMutex      sync.Mutex
	quotaCalcMutex  sync.Mutex
	quotaNotify     chan interface{}
	pendingRequests int
	RawSaveFolder   string
}

// CreateCanvas creates a new Canvas object
func CreateCanvas(subdomain, token string, rawSaveFolder string) (*Canvas, error) {
	c := &Canvas{
		subdomain:      subdomain,
		token:          token,
		parameterTypes: []parameterType{},
		client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		lastQuota:       rateLimitMax,
		lastQuotaTime:   time.Now(),
		quotaMutex:      sync.Mutex{},
		quotaCalcMutex:  sync.Mutex{},
		quotaNotify:     make(chan interface{}),
		pendingRequests: 0,
		RawSaveFolder:   rawSaveFolder,
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
