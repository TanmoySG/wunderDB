package wdbClient

import (
	c "github.com/TanmoySG/wunderDB/internal/collections"
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddCollection(databaseId, collectionId model.Identifier, schema model.Schema) *er.WdbError {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if exists, _ := collections.CheckIfExists(collectionId); exists {
		return &er.CollectionAlreadyExistsError
	}

	collections.CreateCollection(collectionId, schema, model.Metadata{}, model.Access{})
	return nil
}

func (wdb wdbClient) GetCollection(databaseId, collectionId model.Identifier) (*model.Collection, *er.WdbError) {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return nil, &er.CollectionAlreadyExistsError
	}

	return collections.GetCollection(collectionId), nil
}

func (wdb wdbClient) DeleteCollection(databaseId, collectionId model.Identifier) *er.WdbError {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return &er.CollectionDoesNotExistsError
	}

	collections := c.UseDatabase(*database)

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return &er.CollectionAlreadyExistsError
	}

	collections.DeleteCollection(collectionId)
	return nil
}
