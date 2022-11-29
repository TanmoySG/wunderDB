package wdbClient

import (
	"fmt"

	c "github.com/TanmoySG/wunderDB/internal/collections"
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddCollection(databaseId, collectionId model.Identifier, schema model.Schema, metadata model.Metadata) error {
	dbExists, database := wdb.Databases.CheckIfExists(databaseId)
	if !dbExists {
		return fmt.Errorf("error deleting database %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}

	collection := c.UseDatabase(database)

	if exists, _ := collection.CheckIfExists(collectionId); exists {
		return fmt.Errorf("error creating collection: %s", er.CollectionAlreadyExistsError.ErrMessage)
	}
	
	collection.CreateCollection(collectionId, schema, model.Metadata{}, model.Access{})
	return nil
}
