package wdbClient

import (
	"fmt"

	c "github.com/TanmoySG/wunderDB/internal/collections"
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddCollection(databaseId, collectionId model.Identifier, schema model.Schema) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error deleting database %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(database)

	if exists, _ := collections.CheckIfExists(collectionId); exists {
		return fmt.Errorf("error creating collection: %s", er.CollectionAlreadyExistsError.ErrMessage)
	}

	collections.CreateCollection(collectionId, schema, model.Metadata{}, model.Access{})
	return nil
}

func (wdb wdbClient) GetCollection(databaseId, collectionId model.Identifier) (*model.Collection, error) {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return nil, fmt.Errorf("error fetching collection %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(database)

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return nil, fmt.Errorf("error fetching collection %s", er.CollectionAlreadyExistsError.ErrMessage)
	}

	return collections.GetCollection(collectionId), nil
}

func (wdb wdbClient) DeleteCollection(databaseId, collectionId model.Identifier) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error deleting collection: %s", er.CollectionDoesNotExistsError.ErrMessage)
	}

	collections := c.UseDatabase(database)

	if exists, _ := collections.CheckIfExists(collectionId); !exists {
		return fmt.Errorf("error deleting collection %s", er.CollectionAlreadyExistsError.ErrMessage)
	}

	collections.DeleteCollection(collectionId)
	return nil
}
