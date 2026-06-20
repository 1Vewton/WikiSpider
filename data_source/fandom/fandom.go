package fandom

import (
	"github.com/1Vewton/WikiSpider/utils/fandom_requester"
)

// Fandom request data structure
type FandomRequest struct {
	title       string // Title of fandom page
	retry_count int    // The number of retries
	link_limit  int    // The maximum number of references return
	user_agent  string // The user agent
	target_url  string // Target fandom url (MediaWiki api)
}

// Create new fandom request
func New(
	title string,
	retry_count int,
	link_limit int,
	user_agent string,
	target_url string,
) *FandomRequest {
	return &FandomRequest{
		title:       title,
		retry_count: retry_count,
		link_limit:  link_limit,
		user_agent:  user_agent,
		target_url:  target_url,
	}
}

// Get wiki text
func (fandomRequest *FandomRequest) GetText() (string, []string, error) {
	var title string
	texts := make([]string, 0)
	wiki_url := fandom_requester.ConstructWikiTextUrl(
		fandomRequest.target_url,
		fandomRequest.title,
	)
	request_result, err := fandom_requester.GetWikiText(
		wiki_url,
		fandomRequest.user_agent,
		fandomRequest.retry_count,
	)
	if err != nil {
		return title, texts, nil
	}
	title = request_result.Parse.Title
	texts = append(texts, request_result.Parse.WikiText.Content)
	// Default return
	return title, texts, nil
}

// Get wiki references
func (fandomRequest *FandomRequest) GetReferences() ([]string, error) {
	references := make([]string, 0)
	wiki_url := fandom_requester.ConstructWikiReferencesUrl(
		fandomRequest.target_url,
		fandomRequest.title,
		fandomRequest.link_limit,
	)
	request_result, err := fandom_requester.GetWikiReferences(
		wiki_url,
		fandomRequest.user_agent,
		fandomRequest.retry_count,
	)
	if err != nil {
		// Return the error and texts if error occurs
		return references, err
	}
	for _, page := range request_result.Query.Pages {
		for _, link := range page.Links {
			references = append(references, link.Title)
		}
	}
	// Default return
	return references, nil
}

// Get target url
func (fandomRequest *FandomRequest) GetTargetURL() string {
	// Directly return the target url
	return fandomRequest.target_url
}
