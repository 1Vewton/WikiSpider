package wikipedia_requester

import (
	"fmt"
	"net/url"
	"strconv"
)

// Construct url for getting the text from wikipedia page.
func ConstructWikiTextUrl(
	title string,
	target_url string,
) string {
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "query")
	query.Add("titles", title)
	query.Add("prop", "extracts")
	query.Add("explaintext", "true")
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Return the url for getting the texts inside a wiki page.
	return wiki_url
}

// Construct url for getting the references
func ConstructWikiReferencesUrl(
	title string,
	target_url string,
	link_limit int,
) string {
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "query")
	query.Add("generator", "links")
	query.Add("titles", title)
	query.Add("prop", "info")
	query.Add("inprop", "url")
	query.Add("gpllimit", strconv.Itoa(link_limit))
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Return the url for the hyperlinks inside a wiki page.
	return wiki_url
}
