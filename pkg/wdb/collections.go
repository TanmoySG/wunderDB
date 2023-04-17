package wdbClient

import (
	c "github.com/TanmoySG/wunderDB/internal/collections"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

func (wdb wdbClient) AddCollection(databaseId, collectionId model.Identifier, schema model.Schema) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if !wdb.safeName.Check(collectionId.String()) {
		return &er.CollectionNameFormatError
	}

	if exists, _ := collections.CheckIfExists(collectionId); exists {
		return &er.CollectionAlreadyExistsError
	}

	collections.CreateCollection(collectionId, schema, model.Metadata{}, model.Access{})
	return nil
}

func (wdb wdbClient) GetCollection(databaseId, collectionId model.Identifier) (*model.Collection, *er.WdbError) {
	if !wdb.safeName.Check(databaseId.String()) {
		return nil, &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if !wdb.safeName.Check(collectionId.String()) {
		return nil, &er.CollectionNameFormatError
	}

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return nil, &er.CollectionDoesNotExistsError
	}

	return collections.GetCollection(collectionId), nil
}

func (wdb wdbClient) DeleteCollection(databaseId, collectionId model.Identifier) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if !wdb.safeName.Check(collectionId.String()) {
		return &er.CollectionNameFormatError
	}

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return &er.CollectionDoesNotExistsError
	}

	collections.DeleteCollection(collectionId)
	return nil
}
