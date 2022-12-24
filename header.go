package main

import "net/http"

// RequestHeader is a function that adds a header to a request.
type RequestHeader func(req *http.Request)

func AddHeader(key, value string) RequestHeader {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// BasicAuth adds basic auth header to request.
func BasicAuth(login, password string) RequestHeader {
	return func(req *http.Request) {
		if login != "" && password != "" {
			req.SetBasicAuth(login, password)
		}
	}
}
