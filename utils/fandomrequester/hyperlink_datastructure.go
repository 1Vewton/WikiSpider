package fandomrequester

// The data structure for hyperlinks of fandom wiki
type HyperLink struct {
	Ns    int    `json:"ns"`
	Title string `json:"title"`
}

// The data structure for page
type Page struct {
	PageId int         `json:"pageid"`
	Ns     int         `json:"ns"`
	Title  string      `json:"title"`
	Links  []HyperLink `json:"links"`
}

// The data structure for pages
type HyperLinkQuery struct {
	Pages map[string]Page `json:"pages"`
}

// The data structure for the response
type HyperLinkResponse struct {
	Query HyperLinkQuery `json:"query"`
}
