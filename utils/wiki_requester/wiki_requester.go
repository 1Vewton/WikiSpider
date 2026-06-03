package wiki_requester

import (
	"fmt"
	"net/url"

	"github.com/1Vewton/WikiSpider/utils/config"
)

// Get the full text of a wikipedia page
func GetWikiText(title string) {
	// Construct the URL
	wiki_url := config.GetWikiUrl()
	query := url.Values{}
	query.Add("action", "query")
	query.Add("title", title)
	query.Add("prop", "extracts")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
}
