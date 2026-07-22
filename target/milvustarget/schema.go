package milvustarget

import (
	"github.com/milvus-io/milvus/client/v2/entity"
)

// MilvusSchema defines the way of creating schema and upload to the database
type MilvusSchema interface {
	GetSchema() *entity.Schema
}

// DefaultMilvusSchema defines the creation of default schema defined by this program
type DefaultMilvusSchema struct {
}

// CreateDefaultMilvusSchema creates a default milvus schema
func CreateDefaultMilvusSchema() *DefaultMilvusSchema {
	return &DefaultMilvusSchema{}
}

// GetSchema returns the schema by default
func GetSchema() *entity.Schema {
	function := entity.NewFunction().
		WithName("text_bm25_emb").
		WithInputFields("original_text").
		WithOutputFields("text_sparse").
		WithType(entity.FunctionTypeBM25)
	schema := entity.NewSchema()
	schema.WithField(
		// Document id
		entity.NewField().
			WithName("document_id").
			WithDataType(entity.FieldTypeVarChar).
			WithMaxLength(1000).
			WithIsPrimaryKey(true).
			WithIsAutoID(false),
	).WithField(
		// Source URL
		entity.NewField().
			WithName("source_url").
			WithMaxLength(5000).
			WithDataType(entity.FieldTypeString),
	).WithField(
		// Original TExt
		entity.NewField().
			WithName("original_text").
			WithMaxLength(5000).
			WithDataType(entity.FieldTypeString),
	).WithField(
		// Title
		entity.NewField().
			WithName("title").
			WithMaxLength(5000).
			WithDataType(entity.FieldTypeString),
	).WithField(
		// Sparse Vector
		entity.NewField().
			WithName("text_sparce").
			WithDataType(entity.FieldTypeSparseVector),
	).WithField(
		// Dense Vector
		entity.NewField().
			WithName("text_dense").
			WithDataType(entity.FieldTypeFloatVector).
			WithDim(1024),
	).WithFunction(
		function,
	)
	return schema
}
