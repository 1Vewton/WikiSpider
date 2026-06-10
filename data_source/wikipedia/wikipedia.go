package wikipedia

import (
	"github.com/1Vewton/WikiSpider/utils/wikipedia_requester"
)

// Wkipedia request data structure
type WikipediaRequest struct {
	title       string
	retry_count int
	link_limit  int
	user_agent  string
	target_url  string
}

// Create a new WikipediaRequest instance
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

// Get text for a wikipedia page
func (wikipediaRequest WikipediaRequest) GetText() (string, []string, error) {
	var corrected_title string
	request_result, err := wikipedia_requester.GetWikiText(
		wikipediaRequest.target_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.title,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		// When an error occurs, return the error
		return corrected_title, nil, err
	}
	corrected_title = request_result.Query.Normalized[0].To
	process_result := make([]string, len(request_result.Query.Pages))
	for _, page := range request_result.Query.Pages {
		process_result = append(process_result, page.Extract)
	}
	// Return the extracted text
	return corrected_title, process_result, nil
}

// Get references in a wikipedia page
func (wikipediaRequest WikipediaRequest) GetReferences() ([]string, error) {
	request_result, err := wikipedia_requester.GetWikiReferences(
		wikipediaRequest.target_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.title,
		wikipediaRequest.link_limit,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		// When an error occurs, return the error
		return nil, err
	}
	process_result := make([]string, len(request_result.Query.Pages))
	for _, hyperlink := range request_result.Query.Pages {
		process_result = append(process_result, hyperlink.Title)
	}
	// Return the extracted references
	return process_result, nil
}
