package wdbClient

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateRole(roleID model.Identifier, allowedActions []string, deniedActions []string) *er.WdbError {
	if exists, _ := wdb.Roles.CheckIfExists(roleID); exists {
		return &er.RoleAlreadyExistsError
	}
	wdb.Roles.CreateRole(roleID, allowedActions, deniedActions)
	return nil
}

func (wdb wdbClient) ListRole() roles.Roles {
	return wdb.Roles.ListRoles()
}
