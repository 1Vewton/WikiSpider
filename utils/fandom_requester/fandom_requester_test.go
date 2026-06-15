package fandom_requester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const stone_references string = `{
  "continue": {
    "plcontinue": "12|0|Alpha_1.1.0.8",
    "continue": "||"
  },
  "query": {
    "pages": {
      "12": {
        "pageid": 12,
        "ns": 0,
        "title": "Stone",
        "links": [
          {
            "ns": 0,
            "title": "1.0"
          },
          {
            "ns": 0,
            "title": "1.90"
          },
          {
            "ns": 0,
            "title": "Activator Rail"
          },
          {
            "ns": 0,
            "title": "Air"
          },
          {
            "ns": 0,
            "title": "Allow"
          }
        ]
      }
    }
  }
}`
const stone_text string = `{
  "parse": {
    "title": "Stone",
    "pageid": 12,
    "wikitext": {
      "*": "Stone"
	  }
  }
}`

// Test whether the function fetches wiki text successfully
func TestGetWikiText(t *testing.T) {
	// Start the mock server
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(stone_text))
			},
		),
	)
	defer mockServer.Close()
	_, err := GetWikiText(
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
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(stone_text))
			},
		),
	)
	defer mockServer.Close()
	_, err := GetWikiReferences(
		mockServer.URL,
		"modzilla/5.0",
		3,
	)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}
