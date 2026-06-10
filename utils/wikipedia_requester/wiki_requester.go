package wikipedia_requester

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/1Vewton/WikiSpider/utils/common_requester"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

// Get the full text of a wikipedia page
func GetWikiText(
	target_url string,
	user_agent string,
	title string,
	retry_count int,
) (WikiTextResponse, error) {
	var wiki_text_response WikiTextResponse
	// Construct the URL
	wiki_url := target_url
	query := url.Values{}
	query.Add("action", "query")
	query.Add("titles", title)
	query.Add("prop", "extracts")
	query.Add("explaintext", "true")
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	service_logger.Info(
		fmt.Sprintf(
			"Requesting wiki text for url: %s",
			wiki_url,
		),
	)
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

// Get references
func GetWikiReferences(
	target_url string,
	user_agent string,
	title string,
	link_limit int,
	retry_count int,
) (HyperLinkResponse, error) {
	var hyper_link_response HyperLinkResponse
	// Construct the URL
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
	service_logger.Info(
		fmt.Sprintf(
			"Requesting wiki references for url: %s",
			wiki_url,
		),
	)
	// Requesting the wiki references
	text_process_result, err := common_requester.CommonGetFunction(
		wiki_url,
		retry_count,
		user_agent,
	)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error fetching references: %s",
				err,
			),
		)
		// Return error if request fails
		return hyper_link_response, err
	}
	err = json.Unmarshal(text_process_result, &hyper_link_response)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error parsing response json: %s",
				err,
			),
		)
		return hyper_link_response, err
	}
	// Default return
	return hyper_link_response, nil
}
