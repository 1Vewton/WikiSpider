package embedder

import (
	"context"
)

// Embedder defines the basic methods of embedders
type Embedder interface {
	// Embed get the result of embedding
	Embed(ctx context.Context, text string) []float64
	EmbedAll(ctx context.Context, text string) [][]float64
}
