package fandom_requester

import (
	"fmt"
	"net/url"
	"strconv"
)

// Construct the url for getting the wiki text of a page
func ConstructWikiTextUrl(
	target_url string,
	title string,
) string {
	// Construct the url
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "parse")
	query.Add("prop", "wikitext")
	query.Add("format", "json")
	query.Add("page", title)
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Return the url for getting the texts inside a wiki page.
	return wiki_url
}

// Construct the url for getting the references of a page
func ConstructWikiReferencesUrl(
	target_url string,
	title string,
	link_limit int,
) string {
	// Construct the url
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "query")
	query.Add("titles", title)
	query.Add("prop", "links")
	query.Add("pllimit", strconv.Itoa(link_limit))
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Return the url for the hyperlinks inside a wiki page.
	return wiki_url
}
