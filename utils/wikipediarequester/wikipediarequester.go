// Requesting wikipedia
package wikipediarequester

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/commonrequester"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

// Get the full text of a wikipedia page
func GetWikiText(
	wiki_url string,
	user_agent string,
	retry_count int,
) (*WikiTextResponse, error) {
	var wiki_text_response WikiTextResponse
	service_logger.Info(
		fmt.Sprintf(
			"Requesting wiki text for url: %s",
			wiki_url,
		),
	)
	// Requesting the wiki text
	err := commonrequester.CommonGetFunction(
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
		return nil, err
	}
	// Default return
	return &wiki_text_response, nil
}

// Get references
func GetWikiReferences(
	wiki_url string,
	user_agent string,
	retry_count int,
) (*HyperLinkResponse, error) {
	var hyper_link_response HyperLinkResponse
	service_logger.Info(
		fmt.Sprintf(
			"Requesting wiki references for url: %s",
			wiki_url,
		),
	)
	// Requesting the wiki references
	err := commonrequester.CommonGetFunction(
		wiki_url,
		retry_count,
		user_agent,
		&hyper_link_response,
	)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error fetching references: %s",
				err,
			),
		)
		// Return error if request fails
		return nil, err
	}
	// Default return
	return &hyper_link_response, nil
}
