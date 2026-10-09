package milvustarget

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/1Vewton/WikiSpider/utils"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
	"golang.org/x/sync/errgroup"
)

// MilvusTarget defines the target of the milvus
type MilvusTarget struct {
	databaseURL     string
	userName        string
	password        string
	dataBase        string
	collectionName  string
	client          *milvusclient.Client
	primaryField    string
	ClientCloseOnce sync.Once
	createSchema    func(collectionName string) (
		string,
		*entity.Schema,
		[]milvusclient.CreateIndexOption,
	)
	uploadData func(
		ctx context.Context,
		client *milvusclient.Client,
		datas []any,
	) error
}

// CreateMilvusTarget creates a milvus target
func CreateMilvusTarget(
	databaseURL string,
	userName string,
	password string,
	dataBase string,
	collectionName string,
	createSchema func(collectionName string) (
		string,
		*entity.Schema,
		[]milvusclient.CreateIndexOption,
	), uploadData func(
		ctx context.Context,
		client *milvusclient.Client,
		datas []any,
	) error,
) *MilvusTarget {
	return &MilvusTarget{
		databaseURL:    databaseURL,
		userName:       userName,
		password:       password,
		dataBase:       dataBase,
		collectionName: collectionName,
		createSchema:   createSchema,
		uploadData:     uploadData,
	}
}

// Connect connects to the milvus server, timeout is setted
func (target *MilvusTarget) Connect(
	ctx context.Context,
	timeout time.Duration,
) error {
	ctxConnect, cancelConnect := context.WithTimeout(
		ctx,
		timeout,
	)
	defer cancelConnect()
	// Connect to database
	cli, errConn := milvusclient.New(
		ctxConnect,
		&milvusclient.ClientConfig{
			Address:  target.databaseURL,
			Username: target.userName,
			Password: target.password,
		},
	)
	if errConn != nil {
		return errConn
	}
	// Show all databases
	databases, errViewDatabase := cli.ListDatabase(
		ctx,
		milvusclient.NewListDatabaseOption(),
	)
	if errViewDatabase != nil {
		return errViewDatabase
	}
	found := false
	for _, db := range databases {
		if db == target.dataBase {
			found = true
		}
	}
	// Create database
	if !found {
		cli.CreateDatabase(
			ctx,
			milvusclient.NewCreateDatabaseOption(
				target.dataBase,
			),
		)
	}
	// Use database
	err := cli.UseDatabase(
		ctx,
		milvusclient.NewUseDatabaseOption(
			target.dataBase,
		),
	)
	if err != nil {
		return err
	}
	target.client = cli
	return nil
}

// IfValueExists checks if the value exists in the target database
func (target *MilvusTarget) IfValueExists(
	ctx context.Context,
	field string,
	value any,
) (bool, error) {
	queryResult, errQuery := target.client.Query(
		ctx,
		milvusclient.NewQueryOption(
			target.collectionName,
		).WithFilter(
			CreateQueryTextString(
				field,
				value,
			),
		).WithOutputFields(
			target.primaryField,
		),
	)
	if errQuery != nil {
		return false, errQuery
	}
	ids := queryResult.GetColumn(
		target.primaryField,
	).FieldData().GetScalars().GetArrayData().Data
	if len(ids) > 0 {
		return true, nil
	}
	return false, nil
}

// InitializeSchema initializes the schema of the database and creates collection
func (target *MilvusTarget) InitializeSchema(
	ctx context.Context,
) error {
	// List collections
	resCollections, errListCollection := target.client.ListCollections(
		ctx,
		milvusclient.NewListCollectionOption(),
	)
	if errListCollection != nil {
		return errListCollection
	}
	for _, clientName := range resCollections {
		if clientName == target.collectionName {
			return fmt.Errorf(
				"This client %s already exists",
				target.collectionName,
			)
		}
	}
	// Create collection
	primaryField, schema, indexOptions := target.createSchema(
		target.collectionName,
	)
	target.primaryField = primaryField
	errCreateCollection := target.client.CreateCollection(
		ctx,
		milvusclient.NewCreateCollectionOption(
			target.collectionName,
			schema,
		).WithIndexOptions(indexOptions...),
	)
	return errCreateCollection
}

// TmpResult stores the result for checking field
type TmpResult struct {
	Field  string
	Exists bool
}

// CheckFields checks whether certain field has certain value.
// Return true if there are certain field in the database has certain value.
func (target *MilvusTarget) CheckFields(
	ctx context.Context,
	data any,
	fields ...string,
) (bool, error) {
	resultChan := make(chan TmpResult, len(fields))
	group, ctx := errgroup.WithContext(ctx)
	for _, field := range fields {
		tmpField := field
		group.Go(
			func() error {
				resDataInField, errGetField := utils.GetValueFromFieldByString(
					data,
					tmpField,
				)
				if errGetField != nil {
					return errGetField
				}
				result, errIfValueExists := target.IfValueExists(
					ctx,
					tmpField,
					resDataInField,
				)
				if errIfValueExists != nil {
					return errIfValueExists
				}
				resultChan <- TmpResult{
					Field:  tmpField,
					Exists: result,
				}
				return nil
			},
		)
	}
	if err := group.Wait(); err != nil {
		return true, err
	}
	close(resultChan)
	for data := range resultChan {
		if !data.Exists {
			return false, nil
		}
	}
	return true, nil
}

// NewDoc inserts new document into the database
func (target *MilvusTarget) NewDoc(
	ctx context.Context,
	datas []any,
) error {
	err := target.uploadData(
		ctx,
		target.client,
		datas,
	)
	return err
}

// Close closes the connection
func (target *MilvusTarget) Close(
	ctx context.Context,
) {
	target.ClientCloseOnce.Do(
		func() {
			_ = target.client.Close(ctx)
		},
	)
}
