package collections

import (
	"github.com/TanmoySG/wunderDB/model"
)

type Collections map[model.Identifier]*model.Collection

func UseDatabase(database model.Database) Collections {
	return Collections(database.Collections)
}

func (c Collections) CheckIfExists(collectionID model.Identifier) (bool, *model.Collection) {
	collection, collExists := c[collectionID]
	if collExists {
		return collectionExists, collection
	} else {
		return collectionDoesNotExist, collection
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

func (c Collections) GetCollection(collectionID model.Identifier) *model.Collection {
	return c[collectionID]
}

func (c Collections) DeleteCollection(collectionID model.Identifier) {
	delete(c, collectionID)
}
