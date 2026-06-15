package fandom

// Fandom request data structure
type FandomRequest struct {
	title       string // Title of fandom page
	retry_count int    // The number of retries
	link_limit  int    // The maximum number of references return
	user_agent  string // The user agent
	target_url  string // Target fandom url (MediaWiki api)
}
