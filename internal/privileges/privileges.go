package privileges

var (
	Categories = map[string]string{
		// Role Privileges
		CreateRole: RolePrivileges,
		GrantRole:  RolePrivileges,
		UpdateRole: RolePrivileges,

		// Database Privileges
		CreateDatabase: DatabasePrivileges,
		ReadDatabase:   DatabasePrivileges,
		UpdateDatabase: DatabasePrivileges,
		DeleteDatabase: DatabasePrivileges,

		// Collection Privileges
		CreateCollection: CollectionPrivileges,
		ReadCollection:   CollectionPrivileges,
		UpdateCollection: CollectionPrivileges,
		DeleteCollection: CollectionPrivileges,

		// Data Privileges
		AddData:    DataPrivileges,
		ReadData:   DataPrivileges,
		UpdateData: DataPrivileges,
		DeleteData: DataPrivileges,
	}
)

func IsAvailable(privilege string) bool {
	_, privilegeExists := Categories[privilege]
	return privilegeExists
}

func Category(privilege string) string {
	if IsAvailable(privilege) {
		return Categories[privilege]
	}
	return ""
}
