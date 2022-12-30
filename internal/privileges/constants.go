package privileges

const (
	DatabasePrivileges   = "databasePrivileges"
	CollectionPrivileges = "collectionPrivileges"
	DataPrivileges       = "dataPrivileges"
	UserPrivileges       = "usersPrivileges"
	RolePrivileges       = "rolesPrivileges"
)

const (
	Wildcard = "*"
)

const (
	CreateUser = "createUser"

	CreateRole = "createRole"
	GrantRole  = "grantRole"
	UpdateRole = "updateRole"
)

const (
	CreateDatabase = "createDatabase"
	ReadDatabase   = "readDatabase"
	UpdateDatabase = "updateDatabase"
	DeleteDatabase = "deleteDatabase"

	CreateCollection = "createCollection"
	ReadCollection   = "readCollection"
	UpdateCollection = "updateCollection"
	DeleteCollection = "deleteCollection"

	AddData    = "addData"
	ReadData   = "readData"
	UpdateData = "updateData"
	DeleteData = "deleteData"
)

const (
	Allowed = true
	Denied  = false
)
