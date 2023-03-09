package admin

import (
	"encoding/base64"

	p "github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	DEFAULT_ADMIN_USERID          = "admin"
	BASE64_DEFAULT_ADMIN_PASSWORD = "YWRtaW4="

	DEFAULT_ADMIN_ROLE = "wdb_super_admin_role"
)

var (
	ALLOWED_ADMIN_PRIVILEGES = getAllowedRole()
	DENIED_ADMIN_PRIVILEGES  = []string{}
	DEFAULT_ADMIN_PASSWORD   = decodeDefaultPassword(BASE64_DEFAULT_ADMIN_PASSWORD)
	DEFAULT_ADMIN_PERMISSION = getPermission()

	WILDCARD = "*"
)

func getAllowedRole() []string {
	var adminPrivileges []string

	for privilege := range p.PrivilegeScope {
		adminPrivileges = append(adminPrivileges, privilege)
	}

	return adminPrivileges
}

func getPermission() model.Permissions {
	return model.Permissions{
		Role: DEFAULT_ADMIN_ROLE,
		On: &model.Entities{
			Users:       &WILDCARD,
			Databases:   &WILDCARD,
			Collections: &WILDCARD,
		},
	}
}

func decodeDefaultPassword(base64Password string) string {
	data, err := base64.StdEncoding.DecodeString(base64Password)
	if err != nil {
		return ""
	}
	return string(data)
}
