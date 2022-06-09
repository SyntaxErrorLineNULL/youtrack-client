package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type BodyResponse struct {
	Response string `json:"responseData"`
}

type BodyGetResponse []struct {
	Index int    `json:"index"`
	GUID  string `json:"guid"`
}

func TestHTTPClientDoSuccess(t *testing.T) {
	header := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Language": "en",
	}

	client := NewClient(header)

	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "en", r.Header.Get("Accept-Language"))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"responseData":"success"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	var body BodyResponse
	err = client.Request(req.Method, req.URL.String(), req.Body, &body)
	require.NoError(t, err, "failed to make a GET request")

	resJson, err := json.Marshal(body)
	require.NoError(t, err, "[Err] decode json error: %w")

	assert.Equal(t, "{\"responseData\":\"success\"}", string(resJson))
}

func TestHTTPClientGetSuccess(t *testing.T) {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	client := NewClient(header)

	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"index": 0,"guid": "3ebda954-98fb-4254-9b65-1a0e944f393a"}]`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	var body BodyGetResponse
	err = client.Get(req.URL.String(), &body)
	require.NoError(t, err, "failed to make a GET request")

	resJson, err := json.Marshal(body)
	require.NoError(t, err, "[Err] decode json error: %w")
	assert.Equal(t, "[{\"index\":0,\"guid\":\"3ebda954-98fb-4254-9b65-1a0e944f393a\"}]", string(resJson))
}

func TestHTTPClientPostSuccess(t *testing.T) {
	header := map[string]string{
		"Accept":          "application/json",
		"Content-Type":    "application/json",
		"Accept-Language": "en",
	}

	client := NewClient(header)

	requestBodyString := `{ "age": "24" }`

	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "en", r.Header.Get("Accept-Language"))

		rBody, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err, "err extract request body")

		assert.Equal(t, requestBodyString, string(rBody))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{ "responseData": "ok" }`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	requestBody := bytes.NewReader([]byte(requestBodyString))

	var body BodyResponse
	err := client.Post(server.URL, requestBody, &body)
	require.NoError(t, err, "failed to make a POST request")

	resJson, err := json.Marshal(body)
	require.NoError(t, err, "[Err] decode json error: %w")

	assert.Equal(t, "{\"responseData\":\"ok\"}", string(resJson))
}
