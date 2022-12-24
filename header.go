package main

import "net/http"

// RequestHeader is a function that adds a header to a request.
type RequestHeader func(req *http.Request)

func AddHeader(key, value string) RequestHeader {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// BearerAuth adds bearer auth header to request.
func BearerAuth(key string) RequestHeader {
	return func(req *http.Request) {
		if key != "" {
			AddHeader("Authorization", "Bearer "+key)
		}
	}
}
