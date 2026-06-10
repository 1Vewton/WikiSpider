package fandom_requester

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/1Vewton/WikiSpider/utils/common_requester"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

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
	query.Add("prop", "wikitext")
	query.Add("format", "json")
	query.Add("page", title)
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	// Requesting the wiki text
	text_process_result, err := common_requester.CommonGetFunction(
		wiki_url,
		retry_count,
		user_agent,
	)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error requesting wiki text: %s",
				err,
			),
		)
		// Return error if request fails
		return wiki_text_response, err
	}
	err = json.Unmarshal(text_process_result, &wiki_text_response)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error parsing response json: %s",
				err,
			),
		)
		return wiki_text_response, err
	}
	// Default return
	return wiki_text_response, nil
}
