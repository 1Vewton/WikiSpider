package wikipedia_requester

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

// Data structure of hyperlink query
type HyperLinkQuery struct {
	Pages map[string]HyperLink `json:"pages"`
}

// The main structure of the hyperlink response
type HyperLinkResponse struct {
	BatchComplete string         `json:"batchcomplete"`
	Query         HyperLinkQuery `json:"query"`
}
