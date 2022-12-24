package main

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type ClientTransport struct {
	roundTripperWrap http.RoundTripper
	rateLimiter      *rate.Limiter
}

func (c *ClientTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := c.rateLimiter.Wait(r.Context()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	return c.roundTripperWrap.RoundTrip(r)
}

// NewClientTransport wraps transportWrap with a rate limitter
// examle usage:
// client := http.DefaultClient
// client.Transport = NewClientTransport(10*time.Seconds, 60, http.DefaultTransport) allows 60 requests every 10 seconds

func NewClientTransport(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	return &ClientTransport{
		roundTripperWrap: transportWrap,
		rateLimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}
