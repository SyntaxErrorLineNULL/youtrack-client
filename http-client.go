package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type RequestOptions struct {
	Headers []RequestHeader
}

func Get[T any](ctx context.Context, link *url.URL, opts RequestOptions) (T, error) {
	var m T
	request, err := http.NewRequestWithContext(ctx, "GET", link.String(), http.NoBody)
	if err != nil {
		return m, err
	}

	for _, header := range opts.Headers {
		header(request)
	}

	body, err := doRequest(request)
	if err != nil {
		return m, err
	}

	return parseJSON[T](body)
}

func doRequest(r *http.Request) ([]byte, error) {
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case 500:
			// TODO: create error
		}
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseJSON[T any](s []byte) (T, error) {
	var r T
	if err := json.Unmarshal(s, &r); err != nil {
		return r, err
	}
	return r, nil
}
