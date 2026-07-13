package embedder

import (
	"context"
)

// EmbeddingResult stores the result of embedding
type EmbeddingResult struct {
	OriginalText string
	Vector       []float64
}

// Embedder defines the basic methods of embedders
type Embedder interface {
	// Embed get the result of embedding
	Embed(ctx context.Context, text string) ([]float64, error)
	EmbedAll(ctx context.Context, texts []string) ([]EmbeddingResult, error)
}
