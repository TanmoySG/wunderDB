package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	wdbErrors "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/users/admin"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateDefaultAdmin() {
	wdb.Users.CreateUser(admin.DEFAULT_ADMIN_USERID, admin.DEFAULT_ADMIN_PASSWORD, authentication.SHA256, model.Metadata{})
	wdb.Roles.CreateRole(admin.DEFAULT_ADMIN_ROLE, admin.ALLOWED_ADMIN_PRIVILEGES, admin.DENIED_ADMIN_PRIVILEGES)
	wdb.Users.GrantRole(admin.DEFAULT_ADMIN_USERID, admin.DEFAULT_ADMIN_PERMISSION)
}

func (wdb wdbClient) AdminExists() bool {
	adminExists, _ := wdb.Users.CheckIfExists(admin.DEFAULT_ADMIN_USERID)
	return adminExists
}

func (wdb wdbClient) CreateAdmin(adminUserId, password string) *wdbErrors.WdbError {
	error := wdb.CreateUser(model.Identifier(adminUserId), password)
	return error
}

func (wdb wdbClient) HandleAdmin(config config.Config) {
	adminExists, _ := wdb.Users.CheckIfExists(admin.DEFAULT_ADMIN_USERID)
	_ = adminExists
}