package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/internal/users/admin"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
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

	// temporary solution to pick super-admin privileges, instead of latest added - for admin users only.
	// temporary fix for issue https://github.com/TanmoySG/wunderDB/issues/71
	// proper fix will be done in issue https://github.com/TanmoySG/wunderDB/issues/26
	if wdb.Configurations.Admin.String() == userID.String() {
		if wdb.hasAdminPrivileges(userPermissions) {
			return Allowed, nil
		}
	}

	isPermitted := wdb.Roles.Check(userPermissions, privilege, &entities)
	if isPermitted == privileges.Allowed {
		return Allowed, nil
	}
	return Denied, &er.PrivilegeUnauthorized
}

// temporary solution to pick super-admin privileges, instead of latest added - for admin users only.
// temporary fix for issue https://github.com/TanmoySG/wunderDB/issues/71
// proper fix will be done in issue https://github.com/TanmoySG/wunderDB/issues/26
func (wdb wdbClient) hasAdminPrivileges(permissions []model.Permissions) bool {
	for _, permission := range permissions {
		if permission.Role == admin.DEFAULT_ADMIN_ROLE {
			return true
		}
	}
	return false
}
