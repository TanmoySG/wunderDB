package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/roles/sysroles"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

func (wdb wdbClient) AddDatabase(databaseId model.Identifier, userId model.Identifier) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	if exists, _ := wdb.Databases.CheckIfExists(databaseId); exists {
		return &er.DatabaseAlreadyExistsError
	}

	wdb.Databases.CreateDatabase(databaseId, model.Access{})

	return wdb.grantSystemDefaultRole(userId, sysroles.DatabaseAdminRole, databaseId.String())
}

func (wdb wdbClient) GetDatabase(databaseId model.Identifier) (*model.Database, *er.WdbError) {
	if !wdb.safeName.Check(databaseId.String()) {
		return nil, &er.DatabaseNameFormatError
	}

	exists, _ := wdb.Databases.CheckIfExists(databaseId)
	if !exists {
		return nil, &er.DatabaseDoesNotExistsError
	}

	return wdb.Databases.GetDatabase(databaseId), nil
}

func (wdb wdbClient) DeleteDatabase(databaseId model.Identifier) *er.WdbError {
	if !wdb.safeName.Check(databaseId.String()) {
		return &er.DatabaseNameFormatError
	}

	exists, database := wdb.Databases.CheckIfExists(databaseId)
	if !exists {
		return &er.DatabaseDoesNotExistsError
	}

	database.Lock()
	defer database.Unlock()

	wdb.Databases.DeleteDatabase(databaseId)
	return nil
}
