// Wikipedia data source
package wikipedia

import (
	"github.com/1Vewton/WikiSpider/utils/wikipedia_requester"
)

// Wkipedia request data structure
type WikipediaRequest struct {
	title       string // Title of the wikipedia page
	retry_count int    // Number of retries for the request
	link_limit  int    // Maximum number of references to retrieve
	user_agent  string // User agent for the request
	target_url  string // Target URL (wikipedia API URL)
}

// Create new wikipedia request
func New(
	title string,
	retry_count int,
	link_limit int,
	user_agent string,
	target_url string,
) WikipediaRequest {
	// Initialize the WikipediaRequest instance
	return WikipediaRequest{
		title:       title,
		retry_count: retry_count,
		link_limit:  link_limit,
		user_agent:  user_agent,
		target_url:  target_url,
	}
}

// Get text for a wikipedia page.
// Returns the corrected title, extracted text, and any errors.
func (wikipediaRequest WikipediaRequest) GetText() (string, []string, error) {
	var corrected_title string
	wiki_url := wikipedia_requester.ConstructWikiTextUrl(
		wikipediaRequest.title,
		wikipediaRequest.target_url,
	)
	request_result, err := wikipedia_requester.GetWikiText(
		wiki_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		// When an error occurs, return the error
		return corrected_title, nil, err
	}
	if len(request_result.Query.Normalized) > 0 {
		corrected_title = request_result.Query.Normalized[0].To
	} else {
		corrected_title = wikipediaRequest.title
	}
	process_result := make([]string, 0)
	for _, page := range request_result.Query.Pages {
		process_result = append(process_result, page.Extract)
	}
	// Return the extracted text
	return corrected_title, process_result, nil
}

// Get references in a wikipedia page.
// Returns the extracted references and any errors.
func (wikipediaRequest WikipediaRequest) GetReferences() ([]string, error) {
	wiki_url := wikipedia_requester.ConstructWikiReferencesUrl(
		wikipediaRequest.title,
		wikipediaRequest.target_url,
		wikipediaRequest.link_limit,
	)
	request_result, err := wikipedia_requester.GetWikiReferences(
		wiki_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		// When an error occurs, return the error
		return nil, err
	}
	process_result := make([]string, 0)
	for _, hyperlink := range request_result.Query.Pages {
		process_result = append(process_result, hyperlink.Title)
	}
	// Return the extracted references
	return process_result, nil
}

// Get the target url
func (wikipediaRequest WikipediaRequest) GetTargetURL() string {
	// Return the target url that is a private member of WikipediaRequest
	return wikipediaRequest.target_url
}
