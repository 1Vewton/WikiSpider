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

// New request
func New(
	title string,
	retry_count int,
	link_limit int,
	user_agent string,
	target_url string,
) WikipediaRequest {
	return WikipediaRequest{
		title:       title,
		retry_count: retry_count,
		link_limit:  link_limit,
		user_agent:  user_agent,
		target_url:  target_url,
	}
}

// Method: Get text
func (wikipediaRequest WikipediaRequest) GetText() ([]string, error) {
	request_result, err := wikipedia_requester.GetWikiText(
		wikipediaRequest.target_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.title,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		return nil, err
	}
	process_result := make([]string, len(request_result.Query.Pages))
	for _, page := range request_result.Query.Pages {
		process_result = append(process_result, page.Extract)
	}
	return process_result, nil
}

// Method: Get references
func (wikipediaRequest WikipediaRequest) GetReferences() ([]string, error) {
	request_result, err := wikipedia_requester.GetWikiReferences(
		wikipediaRequest.target_url,
		wikipediaRequest.user_agent,
		wikipediaRequest.title,
		wikipediaRequest.link_limit,
		wikipediaRequest.retry_count,
	)
	if err != nil {
		return nil, err
	}
	process_result := make([]string, len(request_result.Query.Pages))
	for _, hyperlink := range request_result.Query.Pages {
		process_result = append(process_result, hyperlink.Title)
	}
	return process_result, nil
}
