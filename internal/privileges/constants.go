package privileges

// Scope of Privilege Type
// TODO: Use for below Privileges (ref. L14-21)
type PrivilegeScopeType string

const (
	UserPrivileges       = "userPrivilege"
	GlobalPrivileges     = "globalPrivilege"
	DatabasePrivileges   = "databasePrivilege"
	CollectionPrivileges = "collectionPrivilege"
)

// Read, Write, Wildcard Action Type
type PrivilegeActionType string

var (
	WildcardPrivilege PrivilegeActionType = "wildcardPrivilege"
	WritePrivilege    PrivilegeActionType = "writePrivilege"
	ReadPrivilege     PrivilegeActionType = "readPrivilege"
)

const (
	Allowed = true
	Denied  = false
)

const (
	Ping     = "ping"
	Wildcard = "*"
)

const (
	CreateUser = "createUser"
	GrantRole  = "grantRole"
	RevokeRole = "revokeRole"
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

	AddRecords    = "addRecords"
	ReadRecords   = "readRecords"
	QueryRecords  = "queryRecords"
	UpdateRecords = "updateRecords"
	DeleteRecords = "deleteRecords"
)
