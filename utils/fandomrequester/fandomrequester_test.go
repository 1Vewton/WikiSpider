package fandomrequester

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test whether the function fetches wiki text successfully
func TestGetWikiText(t *testing.T) {
	// Read testdata
	data, err := os.ReadFile(
		"testdata/stone_text.json",
	)
	if err != nil {
		t.Fatalf("The test failed when reading file due to %s", err)
	}
	// Start the mock server
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
			},
		),
	)
	defer mockServer.Close()
	// Start testing
	_, err = GetWikiText(
		mockServer.URL,
		"modzilla/5.0",
		3,
	)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

// Test whether the function fetches references successfully
func TestGetWikiReferences(t *testing.T) {
	// Read testdata
	data, err := os.ReadFile(
		"testdata/stone_references.json",
	)
	if err != nil {
		t.Fatalf("The test failed when reading file due to %s", err)
	}
	// Start the mockserver
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
			},
		),
	)
	defer mockServer.Close()
	// Start testing
	_, err = GetWikiReferences(
		mockServer.URL,
		"modzilla/5.0",
		3,
	)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}
