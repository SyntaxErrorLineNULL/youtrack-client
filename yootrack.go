package main

import (
	"net/http"
	"sync"
	"time"
)

var (
	once     sync.Once
	client   *http.Client = http.DefaultClient
	endpoint              = ""
)

// Settings yootrack client settings
type Settings struct {
	Token            string
	BaseURL          string
	RateLimitTimeout time.Duration
	RateLimitBursts  int
}

type Client struct {
	token   string
	baseURL string
}

func NewClient(settings *Settings) *Client {
	if settings.BaseURL == "" {
		settings.BaseURL = endpoint
	}

	initClient(settings.RateLimitTimeout, settings.RateLimitBursts)

	return &Client{
		token: "",
	}
}

func initClient(rateLimitTimeout time.Duration, rateLimitBursts int) {
	once.Do(func() {
		client = http.DefaultClient
		if rateLimitTimeout != 0 && rateLimitBursts != 0 {
			client.Transport = NewClientTransport(rateLimitTimeout, rateLimitBursts, http.DefaultTransport)
		}
	})
}
