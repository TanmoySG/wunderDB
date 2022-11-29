package wdbClient

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddDatabase(databaseId string, metadata model.Metadata) error {
	if wdb.Databases.CheckIfExists(model.Identifier(databaseId)) {
		return fmt.Errorf("error creating database %s", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	wdb.Databases.CreateDatabase(model.Identifier(databaseId), metadata, model.Access{})
	return nil
}

func (wdb wdbClient) GetDatabase(databaseId string) (*model.Database, error) {
	if !wdb.Databases.CheckIfExists(model.Identifier(databaseId)) {
		return nil, fmt.Errorf("error creating namespace %s", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	return wdb.Databases.GetDatabase(model.Identifier(databaseId)), nil
}

func (wdb wdbClient) DeleteDatabase(databaseId string) error {
	if wdb.Databases.CheckIfExists(model.Identifier(databaseId)) {
		wdb.Databases.DeleteDatabase(model.Identifier(databaseId))
		return nil
	}
	return fmt.Errorf("error deleting database %s", er.DatabaseDoesNotExistsError.ErrMessage)
}
