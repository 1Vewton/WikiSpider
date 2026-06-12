package fandom_requester

import (
	"fmt"
	"net/url"
)

func ConstructWikiUrl(
	target_url string,
	title string,
) string {
	// construct the url
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "parse")
	query.Add("prop", "wikitext")
	query.Add("format", "json")
	query.Add("page", title)
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	return wiki_url
}
