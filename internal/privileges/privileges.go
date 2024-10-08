package privileges

// TODO: merge PrivilegeScope and PrivilegeType maps
// into on map/struct and use same everywhere, eg:
//
//	map[string]struct{
//		Scope PrivilegeScopeType
//		Type  PrivilegeActionType
//	}
var PrivilegeScope = map[string]string{
	CreateRole:       GlobalPrivileges,
	CreateDatabase:   GlobalPrivileges,
	LoginUser:        GlobalPrivileges,
	ListRole:         GlobalPrivileges,
	GrantRole:        UserPrivileges,
	RevokeRole:       UserPrivileges,
	UpdateRole:       UserPrivileges,
	ReadDatabase:     DatabasePrivileges,
	UpdateDatabase:   DatabasePrivileges,
	DeleteDatabase:   DatabasePrivileges,
	CreateCollection: DatabasePrivileges,
	ReadCollection:   CollectionPrivileges,
	UpdateCollection: CollectionPrivileges,
	DeleteCollection: CollectionPrivileges,
	AddRecords:       CollectionPrivileges,
	ReadRecords:      CollectionPrivileges,
	UpdateRecords:    CollectionPrivileges,
	DeleteRecords:    CollectionPrivileges,
	QueryRecords:     CollectionPrivileges,
}

var PrivilegeType = map[string]PrivilegeActionType{
	CreateRole: WildcardPrivilege,
	LoginUser:  WildcardPrivilege,
	GrantRole:  WildcardPrivilege,
	UpdateRole: WildcardPrivilege,
	ListRole:   WildcardPrivilege,
	RevokeRole: WildcardPrivilege,

	ReadDatabase:   ReadPrivilege,
	ReadCollection: ReadPrivilege,
	ReadRecords:    ReadPrivilege,
	QueryRecords:   ReadPrivilege,

	CreateDatabase:   WritePrivilege,
	UpdateDatabase:   WritePrivilege,
	DeleteDatabase:   WritePrivilege,
	CreateCollection: WritePrivilege,
	UpdateCollection: WritePrivilege,
	DeleteCollection: WritePrivilege,
	AddRecords:       WritePrivilege,
	UpdateRecords:    WritePrivilege,
	DeleteRecords:    WritePrivilege,
}

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

func GetPrivilegeType(privilege string) PrivilegeActionType {
	privilegeType, ok := PrivilegeType[privilege]
	if ok {
		return privilegeType
	}
	return WildcardPrivilege
}
