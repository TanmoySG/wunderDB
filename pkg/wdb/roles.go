package wdbClient

import (
	"fmt"
	"strconv"

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

func (wdb wdbClient) CreateRole(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError {
	if !wdb.safeName.Check(roleID.String()) {
		return &er.EntityNameFormatError
	}

	if exists, _ := wdb.Roles.CheckIfExists(roleID); exists {
		return &er.RoleAlreadyExistsError
	}

	return wdb.Roles.Create(roleID, allowed, denied, hidden)
}

func (wdb wdbClient) UpdateRole(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError {
	if !wdb.safeName.Check(roleID.String()) {
		return &er.EntityNameFormatError
	}

	if exists, _ := wdb.Roles.CheckIfExists(roleID); !exists {
		return &er.RoleDoesNotExistsError
	}

	return wdb.Roles.Update(roleID, allowed, denied, hidden)
}

func (wdb wdbClient) ListRole(requesterId, forceListFlag string) (roles.Roles, *er.WdbError) {
	forceList, err := strconv.ParseBool(forceListFlag)
	if err != nil {
		return nil, &er.WdbError{
			ErrCode:        "encodeDecodeError",
			ErrMessage:     fmt.Sprintf("error parsing force flag: %s", err),
			HttpStatusCode: 406,
		}
	}

	if forceList {
		if requesterId != wdb.Configurations.Admin.String() {
			return nil, &er.WdbError{
				// add to wdb errors pkg
				ErrCode:        "forceActionUnautorized",
				ErrMessage:     fmt.Sprintf("force list all roles not autorized for requester: %s", requesterId),
				HttpStatusCode: 401,
			}
		}
		return wdb.Roles.List(forceList), nil
	}

	// force list : false
	return wdb.Roles.List(forceList), nil
}

func (wdb wdbClient) CheckUserPermissions(userID model.Identifier, privilege string, entities model.Entities) (bool, *er.WdbError) {
	if !wdb.safeName.Check(userID.String()) {
		return Denied, &er.EntityNameFormatError
	}

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
