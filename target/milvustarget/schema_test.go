package milvustarget

import (
	"testing"
)

// TestNewInputDataFromMap tests converting map to input data
func TestNewInputDataFromMap(
	t *testing.T,
) {
	data := map[string]any{
		"DocumentID":   "114514",
		"SourceURL":    "abc",
		"OriginalText": "1123",
		"Title":        "KKK",
		"Dim":          1024,
	}
	encodedMap, err := NewInputDataFromMap(data)
	if err != nil {
		t.Error(err)
	}
	if encodedMap.DocumentID != "114514" {
		t.Errorf(
			"expected %s for document id, got %s",
			"114514",
			encodedMap.DocumentID,
		)
	}
	if encodedMap.Dim != 1024 {
		t.Errorf(
			"expected %d for dimension, got %d",
			encodedMap.Dim,
			1024,
		)
	}
}
