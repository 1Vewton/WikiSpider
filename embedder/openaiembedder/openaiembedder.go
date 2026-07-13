package openaiembedder

import (
	"context"
	"errors"

	"github.com/1Vewton/WikiSpider/utils/logger"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var serviceLogger = logger.NewLogger("Embedder")

// OpenAIEmbedder embeds the text by calling the embedding model through OpenAI API
type OpenAIEmbedder struct {
	client     openai.Client
	model      string
	dimension  int64
	retryCount int
}

// NewOpenAIEmbedder creates a new embedder
func NewOpenAIEmbedder(
	BaseURL string,
	Model string,
	APIKey string,
	dimension int64,
	retryCount int,
) *OpenAIEmbedder {
	return &OpenAIEmbedder{
		client: openai.NewClient(
			option.WithAPIKey(APIKey),
			option.WithBaseURL(APIKey),
		),
		model:      Model,
		dimension:  dimension,
		retryCount: retryCount,
	}
}

// Embed embeds single text
func (embedder *OpenAIEmbedder) Embed(
	ctx context.Context,
	text string,
) ([]float64, error) {
	totalRetries := 0
	var embedResult *openai.CreateEmbeddingResponse
	for totalRetries < embedder.retryCount {
		result, err := embedder.client.Embeddings.New(
			ctx,
			openai.EmbeddingNewParams{
				Input: openai.EmbeddingNewParamsInputUnion{
					OfString: openai.String(text),
				},
				Model:          embedder.model,
				Dimensions:     openai.Int(embedder.dimension),
				EncodingFormat: openai.EmbeddingNewParamsEncodingFormatFloat,
			},
		)
		if err == nil {
			embedResult = result
			break
		} else {
			serviceLogger.Error(err.Error())
			totalRetries += 1
			if totalRetries == 3 {
				return nil, err
			}
		}
	}
	if len(embedResult.Data) < 1 {
		return nil, errors.New("The model returns nothing")
	}
	return embedResult.Data[0].Embedding, nil
}
