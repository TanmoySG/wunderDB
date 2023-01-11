package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/users/admin"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) InitializeAdmin(config *config.Config) {
	var userID, userPassword string
	if wdb.Configurations.Admin == nil {
		if config.AdminID == "" || config.AdminPassword == "" {
			userID = admin.DEFAULT_ADMIN_USERID
			userPassword = admin.DEFAULT_ADMIN_PASSWORD
		} else {
			userID = config.AdminID
			userPassword = authentication.Hash(config.AdminPassword, authentication.SHA256)
		}

		userExists, _ := wdb.Users.CheckIfExists(model.Identifier(userID))
		if !userExists {
			wdb.Users.CreateUser(model.Identifier(userID), userPassword, authentication.SHA256, model.Metadata{})
		}

		wdb.Roles.CreateRole(model.Identifier(admin.DEFAULT_ADMIN_ROLE), admin.ALLOWED_ADMIN_PRIVILEGES, admin.DENIED_ADMIN_PRIVILEGES)
		wdb.Users.GrantRole(model.Identifier(userID), admin.DEFAULT_ADMIN_PERMISSION)
	}
}
