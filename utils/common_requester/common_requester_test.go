package common_requester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the requesting function
func TestRequesting(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected GET method, got %s", r.Method)
				}
				if r.UserAgent() != "Mozilla/5.0" {
					t.Errorf("Expected User-Agent 'Mozilla/5.0', got %s", r.UserAgent())
				}
				// Set the response status code
				w.WriteHeader(http.StatusOK)
				// Write the response body
				w.Write([]byte("Hello, World!"))
			},
		),
	)
	// Start requesting
	result, err := CommonGetFunction(
		mockServer.URL,
		3,
		"Mozilla/5.0",
	)
	// Check for errors
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if string(result) != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", string(result))
	}
	// Close mock server to free up resources
	defer mockServer.Close()
}
