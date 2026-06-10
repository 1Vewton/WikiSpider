package wikipedia_requester

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
