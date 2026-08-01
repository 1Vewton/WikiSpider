package milvustarget

import (
	"context"
	"fmt"
	"time"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

// MilvusTarget defines the target of the milvus
type MilvusTarget struct {
	databaseURL    string
	userName       string
	password       string
	dataBase       string
	collectionName string
	client         *milvusclient.Client
	createSchema   func(collectionName string) (
		*entity.Schema,
		[]milvusclient.CreateIndexOption,
	)
}

// CreateMilvusTarget creates a milvus target
func CreateMilvusTarget(
	databaseURL string,
	userName string,
	password string,
	dataBase string,
	collectionName string,
	createSchema func(collectionName string) (
		*entity.Schema,
		[]milvusclient.CreateIndexOption,
	),
) *MilvusTarget {
	return &MilvusTarget{
		databaseURL:    databaseURL,
		userName:       userName,
		password:       password,
		dataBase:       dataBase,
		collectionName: collectionName,
		createSchema:   createSchema,
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
	schema, indexOptions := target.createSchema(target.collectionName)
	errCreateCollection := target.client.CreateCollection(
		ctx,
		milvusclient.NewCreateCollectionOption(
			target.collectionName,
			schema,
		).WithIndexOptions(indexOptions...),
	)
	return errCreateCollection
}

// NewDoc inserts new document into the database
func NewDoc() error {
	return nil
}
