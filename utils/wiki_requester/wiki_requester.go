package wiki_requester

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/1Vewton/WikiSpider/utils/config"
	"github.com/1Vewton/WikiSpider/utils/construct_ua"
	"github.com/1Vewton/WikiSpider/utils/logger"
)

var service_logger = logger.NewLogger("WikiRequester")

// Data structure of a single page
type Pages struct {
	PageId  int    `json:"pageid"`
	Ns      int    `json:"ns"`
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

// Normalized data structure of title
type Normalized struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Data structure of query
type Query struct {
	Normalized []Normalized     `json:"normalized"`
	Pages      map[string]Pages `json:"pages"`
}

// The main structure of the response
type WikiTextResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

// Get the full text of a wikipedia page
func GetWikiText(title string, retry_count int) (WikiTextResponse, error) {
	var wiki_text_response WikiTextResponse
	// Construct the URL
	wiki_url := config.GetWikiUrl()
	query := url.Values{}
	query.Add("action", "query")
	query.Add("titles", title)
	query.Add("prop", "extracts")
	query.Add("explaintext", "true")
	query.Add("format", "json")
	wiki_url = fmt.Sprintf("%s%s", wiki_url, query.Encode())
	service_logger.Info(fmt.Sprintf("Requesting wiki text for url: %s", wiki_url))
	// Construct the request
	req, err := http.NewRequest("GET", wiki_url, nil)
	if err != nil {
		service_logger.Error(fmt.Sprintf("Error constructing request: %s", err.Error()))
		return wiki_text_response, err
	}
	req.Header.Set("User-Agent", construct_ua.ConstructUA())
	// Start requesting
	client := &http.Client{}
	var resp *http.Response
	for i := 0; i < retry_count; i++ {
		resp, err = client.Do(req)
		if err != nil {
			if i == retry_count-1 {
				service_logger.Error(fmt.Sprintf("Error requesting wiki text: %s", err.Error()))
				return wiki_text_response, err
			} else {
				service_logger.Error(fmt.Sprintf("Error requesting wiki text: %s", err.Error()))
				wait_time := rand.Float64()*2 + 1
				time.Sleep(time.Second * time.Duration(wait_time))
			}
		} else {
			// Process the response body
			service_logger.Info("Reading response body")
			var body []byte
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				service_logger.Error(fmt.Sprintf("Error reading response body: %s", err.Error()))
				return wiki_text_response, err
			} else {
				wiki_text_response = WikiTextResponse{}
				err = json.Unmarshal(body, &wiki_text_response)
				if err != nil {
					service_logger.Error(fmt.Sprintf("Error parsing response body: %s", err.Error()))
					return wiki_text_response, err
				} else {
					service_logger.Info("Response parsed successfully")
				}
			}
			service_logger.Info("Requesting wiki text successful")
			break
		}
	}
	defer resp.Body.Close()
	// Parse the response
	return wiki_text_response, nil
}
