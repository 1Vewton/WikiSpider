package wiki_requester

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/1Vewton/WikiSpider/utils/config"
	"github.com/1Vewton/WikiSpider/utils/construct_ua"
)

// Data structure of a single page
type Pages struct {
	PageId  int    `json:"pageid"`
	Ns      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

// Normalized data structure of title
type Normalized struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Data structure of query
type Query struct {
	Normalized Normalized       `json:"normalized"`
	Pages      map[string]Pages `json:"pages"`
}

// The main structure of the response
type WikiTextResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

// Get the full text of a wikipedia page
func GetWikiText(title string) error {
	// Construct the URL
	wiki_url := config.GetWikiUrl()
	query := url.Values{}
	query.Add("action", "query")
	query.Add("title", title)
	query.Add("prop", "extracts")
	query.Add("explaintext", "true")
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Construct the request
	req, err := http.NewRequest("GET", wiki_url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", construct_ua.ConstructUA())
	return nil
}
