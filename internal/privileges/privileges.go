package privileges

var (
	PrivilegeScope = map[string]string{
		CreateRole:       GlobalPrivileges,
		CreateDatabase:   GlobalPrivileges,
		LoginUser:        GlobalPrivileges,
		ListRole:         GlobalPrivileges,
		GrantRole:        DatabasePrivileges,
		UpdateRole:       DatabasePrivileges,
		ReadDatabase:     DatabasePrivileges,
		UpdateDatabase:   DatabasePrivileges,
		DeleteDatabase:   DatabasePrivileges,
		CreateCollection: DatabasePrivileges,
		ReadCollection:   CollectionPrivileges,
		UpdateCollection: CollectionPrivileges,
		DeleteCollection: CollectionPrivileges,
		AddData:          CollectionPrivileges,
		ReadData:         CollectionPrivileges,
		UpdateData:       CollectionPrivileges,
		DeleteData:       CollectionPrivileges,
	}
)

func IsAvailable(privilege string) bool {
	_, privilegeExists := PrivilegeScope[privilege]
	return privilegeExists
}

func Category(privilege string) string {
	if IsAvailable(privilege) {
		return PrivilegeScope[privilege]
	}
	return ""
}
