package wdbClient

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/model"
)

var (
	Allowed bool = true
	Denied  bool = false
)

func (wdb wdbClient) CreateRole(roleID model.Identifier, allowed []string, denied []string) *er.WdbError {
	if exists, _ := wdb.Roles.CheckIfExists(roleID); exists {
		return &er.RoleAlreadyExistsError
	}
	return wdb.Roles.CreateRole(roleID, allowed, denied)
}

func (wdb wdbClient) ListRole() roles.Roles {
	return wdb.Roles.ListRoles()
}

func (wdb wdbClient) CheckUserPermissions(userID model.Identifier, privilege string, entities model.Entities) (bool, *er.WdbError) {
	if exists, _ := wdb.Users.CheckIfExists(userID); !exists {
		return Denied, &er.UserDoesNotExistError
	}

	userPermissions := wdb.Users.Permission(userID)

	isPermitted := wdb.Roles.Check(userPermissions, privilege, &entities)
	if isPermitted == privileges.Allowed {
		return Allowed, nil
	}
	return Denied, &er.PrivilegeUnauthorized
}
