package databases

import (
	"github.com/TanmoySG/wunderDB/model"
)

const (
	databaseExists       = true
	databaseDoesNotExist = false
)

type Databases map[model.Identifier]*model.Database

func WithWDB(databases Databases) Databases {
	return databases
}

func (d Databases) CheckIfExists(databaseID model.Identifier) (bool, *model.Database) {
	database, dbExists := d[databaseID]
	if dbExists {
		return databaseExists, database
	} else {
		return databaseDoesNotExist, database
	}
}

func (d Databases) CreateDatabase(databaseID model.Identifier, metadata model.Metadata, access model.Access) {
	d[databaseID] = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata:    metadata,
		Access:      map[model.Identifier]*model.Access{},
	}
}

func (d Databases) GetDatabase(databaseID model.Identifier) *model.Database {
	return d[databaseID]
}

func (d Databases) DeleteDatabase(databaseID model.Identifier) {
	delete(d, databaseID)
}
