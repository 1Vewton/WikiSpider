package milvustarget

import (
	"testing"

	"github.com/1Vewton/WikiSpider/target"
)

// Test the generation of query text
func TestQueryTextStringGeneration(
	t *testing.T,
) {
	t.Parallel()
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

// TestInterface tests whether the milvustarget can fit into milvus target
func TestInterface(t *testing.T) {
	t.Parallel()
	var testTarget interface{} = &MilvusTarget{}
	_, ok := testTarget.(target.Target)
	if !ok {
		t.Error("milvus target does not implements methods of target")
	}
}
