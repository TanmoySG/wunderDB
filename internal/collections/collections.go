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

func (c Collections) CheckIfCollectionExists(collectionID model.Identifier) bool {
	_, collExists := c[collectionID]
	if collExists {
		return collectionExists
	} else {
		return collectionDoesNotExist
	}
}

func (c Collections) CreateCollection(collectionID model.Identifier, schema model.Schema, metadata model.Metadata, access model.Access) error {
	if c.CheckIfCollectionExists(collectionID) {
		return fmt.Errorf(CollectionErrorFormat, er.CollectionAlreadyExistsError.ErrCode, "error creating collection", er.CollectionAlreadyExistsError.ErrMessage)
	}
	c[collectionID] = &model.Collection{
		Data:     map[model.Identifier]*model.Datum{},
		Schema:   schema,
		Metadata: metadata,
		Access:   map[model.Identifier]*model.Access{},
	}
	return nil
}

func (c Collections) GetCollection(collectionID model.Identifier) (*model.Collection, error) {
	if !c.CheckIfCollectionExists(collectionID) {
		return nil, fmt.Errorf(CollectionErrorFormat, er.CollectionAlreadyExistsError.ErrCode, "error creating collection", er.CollectionAlreadyExistsError.ErrMessage)
	}
	return c[collectionID], nil
}

func (c Collections) DeleteCollection(collectionID model.Identifier) error {
	if c.CheckIfCollectionExists(collectionID) {
		delete(c, collectionID)
		return nil
	}
	return fmt.Errorf(CollectionErrorFormat, er.CollectionDoesNotExistsError.ErrCode, "error deleting collection", er.CollectionDoesNotExistsError.ErrMessage)
}
