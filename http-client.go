package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	// prepared http client to make service calls
	httpClient *http.Client
	headers    map[string]string
}

type YT struct {
	httpClient *Client
	cfg        *YouTrack
}

func (client *Client) Request(method string, link string, body io.Reader, responseStruct any) error {
	request, err := http.NewRequest(method, link, body)
	if err != nil {
		return fmt.Errorf("new request error: %w", err)
	}

	for key, value := range client.headers {
		request.Header.Set(key, value)
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		// what I will return ?
	}

	bodyResponse, err := io.ReadAll(response.Body)
	resErr := response.Body.Close()
	if resErr != nil {
		return fmt.Errorf("body close error: %w", resErr)
	}
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}

	err = json.Unmarshal(bodyResponse, &responseStruct)
	if err != nil {
		// what I will return ?
	}

	return nil

}
