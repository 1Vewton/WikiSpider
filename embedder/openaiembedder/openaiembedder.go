package openaiembedder

import (
	"context"
	"errors"

	"github.com/1Vewton/WikiSpider/embedder"
	"github.com/1Vewton/WikiSpider/utils/logger"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"golang.org/x/sync/errgroup"
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
func (oaiEmbedder *OpenAIEmbedder) Embed(
	ctx context.Context,
	text string,
) ([]float64, error) {
	totalRetries := 0
	var embedResult *openai.CreateEmbeddingResponse
	for totalRetries < oaiEmbedder.retryCount {
		result, err := oaiEmbedder.client.Embeddings.New(
			ctx,
			openai.EmbeddingNewParams{
				Input: openai.EmbeddingNewParamsInputUnion{
					OfString: openai.String(text),
				},
				Model:          oaiEmbedder.model,
				Dimensions:     openai.Int(oaiEmbedder.dimension),
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

// EmbedAll embeds multiple texts
func (oaiEmbedder *OpenAIEmbedder) EmbedAll(
	ctx context.Context,
	texts []string,
) ([]embedder.EmbeddingResult, error) {
	resultChan := make(chan embedder.EmbeddingResult, len(texts))
	group, ctx := errgroup.WithContext(ctx)
	for _, text := range texts {
		tempText := text
		group.Go(
			func() error {
				result, err := oaiEmbedder.Embed(
					ctx,
					tempText,
				)
				if err != nil {
					return err
				}
				embedResult := embedder.EmbeddingResult{
					OriginalText: tempText,
					Vector:       result,
				}
				select {
				case resultChan <- embedResult:
					return nil
				default:
					return errors.New("The length of the channel is too large")
				}
			},
		)
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	var results []embedder.EmbeddingResult = make(
		[]embedder.EmbeddingResult,
		0,
	)
	close(resultChan)
	for result := range resultChan {
		results = append(results, result)
	}
	return results, nil
}
