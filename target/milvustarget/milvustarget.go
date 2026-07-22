package milvustarget

import (
	"context"
	"time"

	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

// MilvusTarget defines the target of the milvus
type MilvustTarget struct {
	databaseURL    string
	userName       string
	password       string
	dataBase       string
	client         *milvusclient.Client
	databaseSchema *MilvusSchema
}

// CreateMilvusTarget creates a milvus target
func CreateMilvusTarget(
	databaseURL string,
	userName string,
	password string,
	dataBase string,
	databaseSchema *MilvusSchema,
) *MilvustTarget {
	return &MilvustTarget{
		databaseURL:    databaseURL,
		userName:       userName,
		password:       password,
		dataBase:       dataBase,
		databaseSchema: databaseSchema,
	}
}

// Connect connects to the milvus server, timeout is setted
func (target *MilvustTarget) Connect(
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
