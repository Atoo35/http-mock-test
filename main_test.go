package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Atoo35/http-mock-test/http_client"
)

type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, fmt.Errorf("DoFunc is not set")
}
func TestTestCall(t *testing.T) {
	// Create a mock HTTP client
	mockClient := &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Mock the behavior of the HTTP client
			// Return a response or error as needed for your test
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       nil, // You can customize the response body
			}, nil
		},
	}

	// Replace the global HTTP client with the mock client
	http_client.Client = mockClient

	// Call the function you want to test
	TestCall()

	// Add assertions based on the expected behavior of TestCall
}
