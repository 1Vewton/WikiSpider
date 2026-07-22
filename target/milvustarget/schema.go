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
	schema := entity.NewSchema()
	schema.WithField(entity.NewField().WithName("DocumentID").
		WithDataType(entity.FieldTypeString).
		WithIsPrimaryKey(true).
		WithIsAutoID(false),
	)
	schema.WithField(entity.NewField().WithName("Source").
		WithDataType(entity.FieldTypeString),
	)
	return schema
}
