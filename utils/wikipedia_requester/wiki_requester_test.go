package wikipedia_requester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const san_fransisco_text string = `
{
  "batchcomplete": "",
  "query": {
    "normalized": [
      {
        "from": "San_Francisco",
        "to": "San Francisco"
      }
    ],
    "pages": {
      "49728": {
        "pageid": 49728,
        "ns": 0,
		"title": "San Francisco",
		"extract": "..."
	  }
	}
  }
}
`
const san_fransisco_references = `
{
  "batchcomplete": "",
  "continue": {
    "gplcontinue": "49728|0|Alaska_Airlines",
    "continue": "gplcontinue||"
  },
  "query": {
    "pages": {
      "55501010": {
        "pageid": 55501010,
        "ns": 0,
        "title": "10-Minute Walk",
        "contentmodel": "wikitext",
        "pagelanguage": "en",
        "pagelanguagehtmlcode": "en",
        "pagelanguagedir": "ltr",
        "touched": "2026-05-28T18:45:54Z",
        "lastrevid": 1215783050,
        "length": 9192,
        "fullurl": "https://en.wikipedia.org/wiki/10-Minute_Walk",
        "editurl": "https://en.wikipedia.org/w/index.php?title=10-Minute_Walk&action=edit",
        "canonicalurl": "https://en.wikipedia.org/wiki/10-Minute_Walk"
      },
      "73866362": {
        "pageid": 73866362,
        "ns": 0,
        "title": "16th Avenue Tiled Steps",
        "contentmodel": "wikitext",
        "pagelanguage": "en",
        "pagelanguagehtmlcode": "en",
        "pagelanguagedir": "ltr",
        "touched": "2026-06-11T18:48:17Z",
        "lastrevid": 1345370504,
        "length": 8435,
        "fullurl": "https://en.wikipedia.org/wiki/16th_Avenue_Tiled_Steps",
        "editurl": "https://en.wikipedia.org/w/index.php?title=16th_Avenue_Tiled_Steps&action=edit",
        "canonicalurl": "https://en.wikipedia.org/wiki/16th_Avenue_Tiled_Steps"
      }
    }
  }
}
`

// Test getting wiki text
func TestGetWikiText(t *testing.T) {
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(san_fransisco_text))
			},
		),
	)
	_, err := GetWikiText(mockServer.URL, "Modzilla/5.0", 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	defer mockServer.Close()
}

// Test getting wiki references
func TestGetWikiReferences(t *testing.T) {
	mockServer := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(san_fransisco_references))
			},
		),
	)
	_, err := GetWikiReferences(mockServer.URL, "Modzilla/5.0", 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	defer mockServer.Close()
}
