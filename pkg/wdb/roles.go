package wdbClient

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/model"
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
