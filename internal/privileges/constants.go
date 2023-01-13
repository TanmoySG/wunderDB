package privileges

const (
	DatabasePrivileges   = "databasePrivilege"
	CollectionPrivileges = "collectionPrivilege"
	GlobalPrivileges     = "globalPrivilege"
)

const (
	Wildcard = "*"
)

const (
	CreateUser = "createUser"

	CreateRole = "createRole"
	GrantRole  = "grantRole"
	UpdateRole = "updateRole"
	ListRole   = "listRole"
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
