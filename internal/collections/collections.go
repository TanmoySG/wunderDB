package collections

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type Collections map[model.Identifier]*model.Collection

func UseDatabase(database model.Database) Collections {
	return Collections(database.Collections)
}

func (c Collections) CheckIfExists(collectionID model.Identifier) (bool, model.Collection) {
	collection, collExists := c[collectionID]
	if collExists {
		return collectionExists, *collection
	} else {
		return collectionDoesNotExist, *collection
	}
}

func (c Collections) CreateCollection(collectionID model.Identifier, schema model.Schema, metadata model.Metadata, access model.Access) {
	c[collectionID] = &model.Collection{
		Data:     map[model.Identifier]*model.Datum{},
		Schema:   schema,
		Metadata: metadata,
		Access:   map[model.Identifier]*model.Access{},
	}
}

func (c Collections) GetCollection(collectionID model.Identifier) (*model.Collection, error) {
	if exists, _ := c.CheckIfExists(collectionID); !exists {
		return nil, fmt.Errorf(CollectionErrorFormat, er.CollectionAlreadyExistsError.ErrCode, "error creating collection", er.CollectionAlreadyExistsError.ErrMessage)
	}
	return c[collectionID], nil
}

func (c Collections) DeleteCollection(collectionID model.Identifier) error {
	if exists, _ := c.CheckIfExists(collectionID); exists {
		delete(c, collectionID)
		return nil
	}
	return fmt.Errorf(CollectionErrorFormat, er.CollectionDoesNotExistsError.ErrCode, "error deleting collection", er.CollectionDoesNotExistsError.ErrMessage)
}
