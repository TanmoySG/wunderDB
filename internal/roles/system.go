package roles

import "github.com/TanmoySG/wunderDB/internal/privileges"

// system defined roles
type SystemDefaultRole struct {
	RoleID     string
	Privileges []string
}

// default database admin role
var DatabaseAdminRole = SystemDefaultRole{
	RoleID: "database_admin",
	Privileges: []string{
		privileges.GrantRole,
		privileges.ReadDatabase,
		privileges.DeleteDatabase,
		privileges.UpdateDatabase,
		privileges.CreateCollection,
	},
}
