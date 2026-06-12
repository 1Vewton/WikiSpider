// Requester for retrieving information from the given URL
package common_requester

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

// Retrieve information from the given URL using GET method
func CommonGetFunction(
	wiki_url string,
	retry_count int,
	user_agent string,
) ([]byte, error) {
	// Initialize response variable
	var wiki_text_response []byte
	// Constructing request
	req, err := http.NewRequest("GET", wiki_url, nil)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error constructing request: %s",
				err,
			),
		)
		// Return error if request construction fails
		return wiki_text_response, err
	}
	req.Header.Set("User-Agent", user_agent)
	// Start requesting
	client := &http.Client{}
	var resp *http.Response
	for i := 0; i < retry_count; i++ {
		resp, err = client.Do(req)
		if err != nil {
			if i == retry_count-1 {
				service_logger.Error(
					fmt.Sprintf(
						"Error occured when requesting: %s",
						err,
					),
				)
				// Return error if request fails for too many times
				return wiki_text_response, err
			} else {
				// If error occurs, go to next turn
				service_logger.Error(
					fmt.Sprintf(
						"Error occured when requesting: %s",
						err,
					),
				)
				wait_time := rand.Float64()*2 + 1
				service_logger.Info(
					fmt.Sprintf(
						"Waiting for %f seconds before retrying",
						wait_time,
					),
				)
				time.Sleep(time.Second * time.Duration(wait_time))
			}
		} else {
			// Process the response body
			service_logger.Info("Reading response body")
			var body []byte
			body, err = io.ReadAll(resp.Body)
			wiki_text_response = body
			service_logger.Info(string(wiki_text_response))
			if err != nil {
				service_logger.Error(
					fmt.Sprintf(
						"Error reading response body: %s",
						err,
					),
				)
				// Return error if response processing fails
				return wiki_text_response, err
			}
			service_logger.Info("Requesting successful")
			break
		}
	}
	// Close response to free up resources
	defer resp.Body.Close()
	// Return the response if no error occurs
	return wiki_text_response, nil
}
