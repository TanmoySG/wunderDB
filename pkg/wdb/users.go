package wdbClient

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateUser(userID model.Identifier, password string) *er.WdbError {
	if exists, _ := wdb.Roles.CheckIfExists(userID); exists {
		return &er.RoleAlreadyExistsError
	}
	wdb.Users.CreateUser(userID, password, wdb.HashingAlgorithm, model.Metadata{})
	return nil
}

func (wdb wdbClient) GrantRoles(userID model.Identifier, permissions []model.Permissions) *er.WdbError {
	if exists, _ := wdb.Roles.CheckIfExists(userID); exists {
		return &er.RoleAlreadyExistsError
	}
	wdb.Users.GrantRole(userID, permissions)
	return nil
}
