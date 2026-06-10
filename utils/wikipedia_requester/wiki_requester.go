package wikipedia_requester

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

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

// Data structure of hyperlink query
type HyperLinkQuery struct {
	Pages map[string]HyperLink `json:"pages"`
}

// Hyper Link
type HyperLink struct {
	PageID           int    `json:"pageid"`
	Ns               int    `json:"ns"`
	Title            string `json:"title"`
	ContentModel     string `json:"contentmodel"`
	PageLanguage     string `json:"pagelanguage"`
	PageLanguageHTML string `json:"pagelanguagehtmlcode"`
	PageLanguageDir  string `json:"pagelanguagedir"`
	Touched          string `json:"touched"`
	LastRevID        int    `json:"lastrevid"`
	Length           int    `json:"length"`
	FullURL          string `json:"fullurl"`
	EditURL          string `json:"editurl"`
	CanonicalURL     string `json:"canonicalurl"`
}

// The main structure of the response
type WikiTextResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

// The main structure of the hyperlink response
type HyperLinkResponse struct {
	BatchComplete string         `json:"batchcomplete"`
	Query         HyperLinkQuery `json:"query"`
}

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
	// Construct the request
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
						"Error requesting wiki text: %s",
						err,
					),
				)
				// Return error if request fails
				return wiki_text_response, err
			} else {
				service_logger.Error(
					fmt.Sprintf(
						"Error requesting wiki text: %s",
						err,
					),
				)
				wait_time := rand.Float64()*2 + 1
				time.Sleep(time.Second * time.Duration(wait_time))
			}
		} else {
			// Process the response body
			service_logger.Info("Reading response body")
			var body []byte
			body, err = io.ReadAll(resp.Body)
			service_logger.Info(string(body))
			if err != nil {
				service_logger.Error(
					fmt.Sprintf(
						"Error reading response body: %s",
						err,
					),
				)
				// Return error if response processing fails
				return wiki_text_response, err
			} else {
				wiki_text_response = WikiTextResponse{}
				err = json.Unmarshal(body, &wiki_text_response)
				if err != nil {
					service_logger.Error(
						fmt.Sprintf(
							"Error parsing response body: %s",
							err,
						),
					)
					// Return error if response parsing fails
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
	// Return the response if no error occurs
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
	// Construct the request
	req, err := http.NewRequest("GET", wiki_url, nil)
	if err != nil {
		service_logger.Error(
			fmt.Sprintf(
				"Error constructing request: %s",
				err,
			),
		)
		// Return error if request construction fails
		return hyper_link_response, err
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
					fmt.Sprintf("Error requesting wiki text: %s",
						err,
					),
				)
				// Return error if request fails
				return hyper_link_response, err
			} else {
				service_logger.Error(
					fmt.Sprintf(
						"Error requesting wiki text: %s",
						err,
					),
				)
				wait_time := rand.Float64()*2 + 1
				time.Sleep(time.Second * time.Duration(wait_time))
			}
		} else {
			// Process the response body
			service_logger.Info("Reading response body")
			var body []byte
			body, err = io.ReadAll(resp.Body)
			service_logger.Info(string(body))
			if err != nil {
				service_logger.Error(
					fmt.Sprintf("Error reading response body: %s",
						err,
					),
				)
				// Return error if response processing fails
				return hyper_link_response, err
			} else {
				hyper_link_response = HyperLinkResponse{}
				err = json.Unmarshal(body, &hyper_link_response)
				if err != nil {
					service_logger.Error(
						fmt.Sprintf(
							"Error parsing response body: %s",
							err,
						),
					)
					// Return error if response parsing fails
					return hyper_link_response, err
				} else {
					service_logger.Info("Response parsed successfully")
				}
			}
			service_logger.Info("Requesting wiki text successful")
			break
		}
	}
	defer resp.Body.Close()
	// Return the response if no error occurs
	return hyper_link_response, nil
}
