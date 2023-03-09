package privileges

const (
	UserPrivileges       = "userPrivilege"
	GlobalPrivileges     = "globalPrivilege"
	DatabasePrivileges   = "databasePrivilege"
	CollectionPrivileges = "collectionPrivilege"
)

const (
	Wildcard = "*"
)

const (
	CreateUser = "createUser"
	GrantRole  = "grantRole"
	LoginUser  = "loginUser"

	CreateRole = "createRole"
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
