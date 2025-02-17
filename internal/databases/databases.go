package databases

import (
	"github.com/TanmoySG/wunderDB/internal/metadata"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/model/redacted"
)

const (
	databaseExists       = true
	databaseDoesNotExist = false
)

type Databases map[model.Identifier]*model.Database

func From(databases Databases) Databases {
	return databases
}

func (d Databases) CheckIfExists(databaseID model.Identifier) (bool, *model.Database) {
	database, dbExists := d[databaseID]
	if dbExists {
		return databaseExists, database
	}

	return databaseDoesNotExist, database
}

func (d Databases) CreateDatabase(databaseID model.Identifier, access model.Access) {
	d[databaseID] = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Access:      map[model.Identifier]*model.Access{},
		Metadata:    metadata.New().BasicChangeMetadata(),
	}
}

func (d Databases) GetDatabase(databaseID model.Identifier) *redacted.RedactedD {
	db := d[databaseID]
	if db == nil {
		return nil
	}

	redactedDb := redacted.RedactedD{Collections: []model.Identifier{}, Metadata: db.Metadata, Access: db.Access}

	for collectionID := range db.Collections {
		redactedDb.Collections = append(redactedDb.Collections, collectionID)
	}

	return &redactedDb
}

func (d Databases) DeleteDatabase(databaseID model.Identifier) {
	delete(d, databaseID)
}

func (d Databases) UpdateMetadata(databaseID model.Identifier) {
	d[databaseID].Metadata = metadata.Use(d[databaseID].Metadata).BasicChangeMetadata()
}
