package milvustarget

import (
	"testing"
)

// Test the generation of query text
func TestQueryTextStringGeneration(
	t *testing.T,
) {
	QueryColor := CreateQueryTextString(
		"color",
		"red",
	)
	if QueryColor != "color like \"red%\"" {
		t.Errorf(
			"Generation failed as the generated result is %s",
			QueryColor,
		)
	}
}
