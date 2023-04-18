package databases

import (
	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	databaseExists       = true
	databaseDoesNotExist = false
)

type Databases map[model.Identifier]*model.Database

func Use(databases Databases) Databases {
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

func (d Databases) CreateDatabase(databaseID model.Identifier, access model.Access) {
	d[databaseID] = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Access:      map[model.Identifier]*model.Access{},
		Metadata:    metadata.New().BasicChangeMetadata(),
	}
}

func (d Databases) GetDatabase(databaseID model.Identifier) *model.Database {
	return d[databaseID]
}

func (d Databases) DeleteDatabase(databaseID model.Identifier) {
	delete(d, databaseID)
}

func (d Databases) UpdateMetadata(databaseID model.Identifier) {
	d[databaseID].Metadata = metadata.Use(d[databaseID].Metadata).BasicChangeMetadata()
}
