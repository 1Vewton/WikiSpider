package openaiembedder

import (
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// OpenAIEmbedder embeds the text by calling the embedding model through OpenAI API
type OpenAIEmbedder struct {
	client openai.Client
	model  string
}

func NewOpenAIEmbedder(
	BaseURL string,
	Model string,
	APIKey string,
) *OpenAIEmbedder {
	return &OpenAIEmbedder{
		client: openai.NewClient(
			option.WithAPIKey(APIKey),
			option.WithBaseURL(APIKey),
		),
		model: Model,
	}
}
