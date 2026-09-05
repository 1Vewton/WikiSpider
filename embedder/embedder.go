package embedder

import (
	"context"
)

// EmbeddingResult stores the result of embedding
type EmbeddingResult struct {
	OriginalText string
	Vector       []float64
}

// Vector64To32 converts embedding result's vector to float32
func (res *EmbeddingResult) Vector64To32() []float32 {
	newVector := make([]float32, len(res.Vector))
	for i, element := range newVector {
		newVector[i] = float32(element)
	}
	return newVector
}

// Embedder defines the basic methods of embedders
type Embedder interface {
	// Embed get the result of embedding
	Embed(ctx context.Context, text string) ([]float64, error)
	EmbedAll(ctx context.Context, texts []string) ([]EmbeddingResult, error)
}
