package canvas

import (
	"fmt"
	"net/http"
	"time"
)

const (
	rateLimitMax       = 700
	rateLimitOutflow   = 10 // Hz
	rateLimitPreflight = 50
)

func (c *Canvas) onRequestStart() {
	c.quotaMutex.Lock()
	defer c.quotaMutex.Unlock()
	for quota := c.GetQuotaAvailable(); quota < rateLimitPreflight; {
		needed := (rateLimitPreflight - quota) / rateLimitOutflow
		timer := time.NewTimer(time.Duration(needed * float64(time.Second)))
		select {
		case <-timer.C:
			break
		case <-c.quotaNotify:
			break
		}
		timer.Stop()
	}
	c.pendingRequests++
}

func (c *Canvas) onRequestFinish(res *http.Response, err error) {
	c.pendingRequests--
	now := time.Now()
	header := res.Header.Get("X-Rate-Limit-Remaining")
	if len(header) > 0 {
		var rem float64
		fmt.Sscanf(header, "%f", &rem)
		c.quotaCalcMutex.Lock()
		defer c.quotaCalcMutex.Unlock()
		c.lastQuota = rem
		c.lastQuotaTime = now
	}
}

// GetQuotaAvailable gets the amount of quota that is currently available
func (c *Canvas) GetQuotaAvailable() float64 {
	c.quotaCalcMutex.Lock()
	defer c.quotaCalcMutex.Unlock()
	dt := time.Now().Sub(c.lastQuotaTime)
	tot := c.lastQuota + rateLimitOutflow*dt.Seconds()
	if tot >= rateLimitMax {
		return rateLimitMax
	}
	return tot - float64(rateLimitPreflight*c.pendingRequests)
}
