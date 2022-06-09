package main

import "net/http"

type Client struct {
	// prepared http client to make service calls
	httpClient *http.Client
	headers    map[string]string
}

type YT struct {
	httpClient *Client
	cfg        *YouTrack
}
