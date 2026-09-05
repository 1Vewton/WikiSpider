package milvustarget

import (
	"context"
	"errors"

	"github.com/1Vewton/WikiSpider/embedder/openaiembedder"
	"github.com/google/uuid"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/mitchellh/mapstructure"
)

// InputData defines the data to input into the schema.
// This is the default schema.
type InputData struct {
	DocumentID   string
	SourceURL    string
	OriginalText string
	Title        string
	TextDense    []float32
	Dim          int
}

// NewInputData creates new input data
func NewInputData(
	documentID string,
	sourceURL string,
	originalText string,
	title string,
	textDense []float32,
) *InputData {
	return &InputData{
		DocumentID:   documentID,
		SourceURL:    sourceURL,
		OriginalText: originalText,
		Title:        title,
		TextDense:    textDense,
		Dim:          len(textDense),
	}
}

// NewInputDataFromMap creates new input data from map
func NewInputDataFromMap(
	data map[string]any,
) (*InputData, error) {
	var result *InputData
	err := mapstructure.Decode(
		data,
		&result,
	)
	return result, err
}

// StringToInputData converts string to input data
func StringToInputData(
	ctx context.Context,
	embedder *openaiembedder.OpenAIEmbedder,
	texts []string,
	sourceURL string,
	title string,
) ([]*InputData, error) {
	result := []*InputData{}
	embeddedChunks, err := embedder.EmbedAll(
		ctx,
		texts,
	)
	if err != nil {
		return nil, err
	}
	for _, embeddedChunk := range embeddedChunks {
		tmpID := uuid.NewString()
		singleInputData := NewInputData(
			tmpID,
			sourceURL,
			embeddedChunk.OriginalText,
			title,
			embeddedChunk.Vector64To32(),
		)
		result = append(result, singleInputData)
	}
	return result, nil
}

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

// GetDefaultSchema returns the schema by default
func GetDefaultSchema(collectionName string) (
	string,
	*entity.Schema,
	[]milvusclient.CreateIndexOption,
) {
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
			WithDataType(entity.FieldTypeVarChar),
	).WithField(
		// Original TExt
		entity.NewField().
			WithName("original_text").
			WithMaxLength(5000).
			WithDataType(entity.FieldTypeVarChar),
	).WithField(
		// Title
		entity.NewField().
			WithName("title").
			WithMaxLength(5000).
			WithDataType(entity.FieldTypeVarChar),
	).WithField(
		// Sparse Vector
		entity.NewField().
			WithName("text_sparse").
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
	indexOptions := []milvusclient.CreateIndexOption{
		milvusclient.NewCreateIndexOption(
			collectionName,
			"text_dense",
			index.NewAutoIndex(
				index.MetricType(entity.IP),
			),
		),
		milvusclient.NewCreateIndexOption(
			collectionName,
			"text_sparse",
			index.NewSparseInvertedIndex(
				entity.BM25,
				0.2,
			),
		),
	}
	return "document_id", schema, indexOptions
}

// UploadData uploads the data
func UploadData(
	ctx context.Context,
	client *milvusclient.Client,
	rawDatas any,
) error {
	datas, ok := rawDatas.([]*InputData)
	if !ok {
		return errors.New(
			"this data is not a list of InputData",
		)
	}
	if len(datas) < 1 {
		return errors.New(
			"data array cannot have length of 0",
		)
	}
	dimension := datas[0].Dim
	// datas
	documentIDArray := []string{}
	sourceURLArray := []string{}
	originalTextArray := []string{}
	titleArray := []string{}
	textDenseArray := [][]float32{}
	for _, data := range datas {
		documentIDArray = append(documentIDArray, data.DocumentID)
		sourceURLArray = append(sourceURLArray, data.SourceURL)
		originalTextArray = append(originalTextArray, data.OriginalText)
		titleArray = append(titleArray, data.Title)
		textDenseArray = append(textDenseArray, data.TextDense)
	}
	_, err := client.Insert(
		ctx,
		milvusclient.NewColumnBasedInsertOption("quick_setup").
			WithVarcharColumn("document_id", documentIDArray).
			WithVarcharColumn("source_url", sourceURLArray).
			WithVarcharColumn("original_text", originalTextArray).
			WithVarcharColumn("title", titleArray).
			WithFloatVectorColumn(
				"text_dense",
				dimension,
				textDenseArray,
			),
	)
	return err
}
