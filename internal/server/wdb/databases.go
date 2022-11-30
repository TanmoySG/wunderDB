package wdbClient

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) AddDatabase(databaseId model.Identifier) error {
	if exists, _ := wdb.Databases.CheckIfExists(databaseId); exists {
		return fmt.Errorf("error creating database %s", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	wdb.Databases.CreateDatabase(databaseId, model.Metadata{}, model.Access{})
	return nil
}

func (wdb wdbClient) GetDatabase(databaseId model.Identifier) (*model.Database, error) {
	if exists, _ := wdb.Databases.CheckIfExists(databaseId); !exists {
		return nil, fmt.Errorf("error creating namespace %s", er.DatabaseAlreadyExistsError.ErrMessage)
	}
	return wdb.Databases.GetDatabase(databaseId), nil
}

func (wdb wdbClient) DeleteDatabase(databaseId model.Identifier) error {
	if exists, _ := wdb.Databases.CheckIfExists(databaseId); !exists {
		return fmt.Errorf("error deleting database %s", er.DatabaseDoesNotExistsError.ErrMessage)
	}
	wdb.Databases.DeleteDatabase(databaseId)
	return nil
}
