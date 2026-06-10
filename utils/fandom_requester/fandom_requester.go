package fandom_requester

import (
	"fmt"
	"net/url"
)

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
