package fandom_requester

import (
	"fmt"
	"net/url"
)

// Wiki Text
type WikiTextStruct struct {
	Content string `json:"*"`
}

// The parse
type ParseStruct struct {
	Title          string         `json:"title"`
	PageID         int            `json:"pageid"`
	WikiTextStruct WikiTextStruct `json:"wikitext"`
}

// The response fetched from the fandom wiki api
type WikiTextResponse struct {
	Parse ParseStruct `json:"parse"`
}

// Get the wiki text
func GetWikiText(
	target_url string,
	user_agent string,
	title string,
	retry_count int,
) (WikiTextResponse, error) {
	var wiki_text_response WikiTextResponse
	// construct the url
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "parse")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Default returning value
	return wiki_text_response, nil
}
