package wikipedia

import (
	"github.com/1Vewton/WikiSpider/utils/wikipedia_requester"
)

// Wkipedia request data structure
type WikipediaRequest struct {
	title       string
	retry_count int
}

// Method: Get text
func (wikipediaRequest WikipediaRequest) GetText() ([]string, error) {
	request_result, err := wikipedia_requester.GetWikiText(wikipediaRequest.title, wikipediaRequest.retry_count)
	if err != nil {
		return nil, err
	}
	process_result := make([]string, len(request_result.Query.Pages))
	for _, page := range request_result.Query.Pages {
		process_result = append(process_result, page.Extract)
	}
	return process_result, nil
}
