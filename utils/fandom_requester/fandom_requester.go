// Requesting fandom
package fandom_requester

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/common_requester"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

// Get the wiki text
func GetWikiText(
	wiki_url string,
	user_agent string,
	retry_count int,
) (WikiTextResponse, error) {
	var wiki_text_response WikiTextResponse
	// Requesting the wiki text
	err := common_requester.CommonGetFunction(
		wiki_url,
		retry_count,
		user_agent,
		&wiki_text_response,
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
	// Default return
	return wiki_text_response, nil
}
