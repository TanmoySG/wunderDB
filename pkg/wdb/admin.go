package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/users/admin"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
)

// TODO: handle individual errors from methods. eg: line 48
func (wdb wdbClient) InitializeAdmin(config *config.Config) {
	var userID, userPassword string
	if config.AdminID == "" {
		userID = admin.DEFAULT_ADMIN_USERID
		userPassword = authentication.Hash(config.AdminPassword, authentication.SHA256)
	} else {
		userID = config.AdminID
		userPassword = authentication.Hash(config.AdminPassword, authentication.SHA256)
	}
	wdb.createAdminRole()
	wdb.processAdmin(userID, userPassword)
}

func (wdb wdbClient) processAdmin(userID, userPassword string) {
	userExists, userDetails := wdb.Users.CheckIfExists(model.Identifier(userID))
	if !userExists {
		wdb.Users.CreateUser(model.Identifier(userID), userPassword, authentication.SHA256, model.Metadata{})
	}

	// check if admin already has access to super admin role
	hasSuperAdminAccess := false

	if userDetails != nil && userDetails.Permissions != nil {
		for _, permission := range userDetails.Permissions {
			if permission.Role == admin.DEFAULT_ADMIN_PERMISSION.Role {
				hasSuperAdminAccess = true
				break
			}
		}
	}

	// grant access to super admin role if not exists
	if !hasSuperAdminAccess {
		wdb.Users.GrantRole(model.Identifier(userID), admin.DEFAULT_ADMIN_PERMISSION)
	}
}

func (wdb wdbClient) createAdminRole() {
	roleExists, _ := wdb.Roles.CheckIfExists(model.Identifier(admin.DEFAULT_ADMIN_ROLE))
	if !roleExists {
		_ = wdb.Roles.CreateRole(model.Identifier(admin.DEFAULT_ADMIN_ROLE), admin.ALLOWED_ADMIN_PRIVILEGES, admin.DENIED_ADMIN_PRIVILEGES, admin.DEFAULT_ADMIN_ROLE_HIDDEN_STATUS)
	}
}
