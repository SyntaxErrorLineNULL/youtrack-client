package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type BodyResponse struct {
	Response string `json:"responseData"`
}

func TestGetSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.SetBasicAuth("customer_key", "customer_secret")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"responseData":"success"}`))
		assert.Equal(t, "Basic Y3VzdG9tZXJfa2V5OmN1c3RvbWVyX3NlY3JldA==", r.Header.Get("Authorization"))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, bytes.NewReader([]byte(`{"message":"Invalid authentication credentials"}`)))
	require.NoError(t, err)

	ctx := context.Background()
	opts := RequestOptions{
		Headers: []RequestHeader{
			AddHeader("Content-Type", "application/json"),
			AddHeader("Accept-Language", "en"),
			BasicAuth("customer_key", "customer_secret"),
		},
	}

	body, err := Get[BodyResponse](ctx, req.URL, opts)
	require.NoError(t, err, "appid is invalid")
	require.NoError(t, err, "invalid authentication credentials")
	require.NoError(t, err, "failed request")

	require.Equal(t, "success", body.Response)
}

func TestGetFailed(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"Invalid authentication credentials"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, bytes.NewReader([]byte(`{"message":"Invalid authentication credentials"}`)))
	require.NoError(t, err)

	ctx := context.Background()
	opts := RequestOptions{
		Headers: []RequestHeader{
			AddHeader("Content-Type", "application/json"),
			AddHeader("Accept-Language", "en"),
			BasicAuth("customer_key", "customer_secret"),
		},
	}

	body, _ := Get[BodyResponse](ctx, req.URL, opts)
	require.Equal(t, "", body.Response)
}
