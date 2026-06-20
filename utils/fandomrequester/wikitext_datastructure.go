package fandomrequester

// Wiki Text
type WikiTextStruct struct {
	Content string `json:"*"`
}

// The parse
type ParseStruct struct {
	Title    string         `json:"title"`
	PageID   int            `json:"pageid"`
	WikiText WikiTextStruct `json:"wikitext"`
}

// The response fetched from the fandom wiki api
type WikiTextResponse struct {
	Parse ParseStruct `json:"parse"`
}
