package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/users/admin"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateDeafultAdmin() {
	wdb.Users.CreateUser(admin.DEFAULT_ADMIN_USERID, admin.DEFAULT_ADMIN_PASSWORD, authentication.SHA256, model.Metadata{})
	wdb.Roles.CreateRole(admin.DEFAULT_ADMIN_ROLE, admin.ALLOWED_ADMIN_PRIVILEGES, admin.DENIED_ADMIN_PRIVILEGES)
	wdb.Users.GrantRole(admin.DEFAULT_ADMIN_USERID, admin.DEFAULT_ADMIN_PERMISSION)
}
