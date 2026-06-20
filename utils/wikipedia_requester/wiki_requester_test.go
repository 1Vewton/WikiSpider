package wikipedia_requester

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test getting wiki text
func TestGetWikiText(t *testing.T) {
	// Read testdata
	data, err := os.ReadFile(
		"testdata/san_fransisco_text.json",
	)
	if err != nil {
		t.Fatalf("The test failed when reading file due to %s", err)
	}
	//Start server
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
	_, err = GetWikiText(mockServer.URL, "Modzilla/5.0", 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test getting wiki references
func TestGetWikiReferences(t *testing.T) {
	// Read testdata
	data, err := os.ReadFile(
		"testdata/san_fransisco_references.json",
	)
	if err != nil {
		t.Fatalf("The test failed when reading file due to %s", err)
	}
	// Start server
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
	_, err = GetWikiReferences(mockServer.URL, "Modzilla/5.0", 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
