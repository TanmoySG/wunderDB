package databases

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type Databases map[model.Identifier]*model.Database

func UseDatabases(wdb model.WDB) Databases {
	return wdb.Databases
}

func (d Databases) CheckIfDatabaseExists(databaseID model.Identifier) bool {
	_, dbExists := d[databaseID]
	if dbExists {
		return databaseExists
	} else {
		return databaseDoesNotExist
	}
}

func (d Databases) CreateNewDatabase(databaseID model.Identifier, metadata model.Metadata, access model.Access) error {
	if d.CheckIfDatabaseExists(databaseID) {
		return fmt.Errorf(DatabaseErrorFormat, er.DatabaseAlreadyExistsError.ErrCode, "error creating database", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	d[databaseID] = &model.Database{
		Collections: map[model.Identifier]*model.Collection{},
		Metadata:    metadata,
		Access:      map[model.Identifier]*model.Access{},
	}
	return nil
}

func (d Databases) GetNamespace(databaseID model.Identifier) (*model.Database, error) {
	if !d.CheckIfDatabaseExists(databaseID) {
		return nil, fmt.Errorf(DatabaseErrorFormat, er.DatabaseAlreadyExistsError.ErrCode, "error creating namespace", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	return d[databaseID], nil
}

func (d Databases) DeleteDatabase(databaseID model.Identifier) error {
	if d.CheckIfDatabaseExists(databaseID) {
		delete(d, databaseID)
		return nil
	}
	return fmt.Errorf(DatabaseErrorFormat, er.DatabaseDoesNotExistsError.ErrCode, "error deleting database", er.DatabaseDoesNotExistsError.ErrMessage)
}
