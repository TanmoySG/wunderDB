package admin

import (
	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	DEFAULT_ADMIN_USERID   = "admin"
	DEFAULT_ADMIN_PASSWORD = "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"

	DEFAULT_ADMIN_ROLE = "wdb_super_admin_role"
)

var (
	ALLOWED_ADMIN_PRIVILEGES = getAllowedRole()
	DENIED_ADMIN_PRIVILEGES  = []string{}

	DEFAULT_ADMIN_PERMISSION = getPermission()

	WILDCARD = "*"
)

func getAllowedRole() []string {
	var adminPrivileges []string

	for privilege, _ := range p.PrivilegeScope {
		adminPrivileges = append(adminPrivileges, privilege)
	}

	return adminPrivileges
}

func getPermission() model.Permissions {
	return model.Permissions{
		Role: DEFAULT_ADMIN_ROLE,
		On: &model.Entities{
			Databases:   &WILDCARD,
			Collections: &WILDCARD,
		},
	}
}
