package privileges

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

	CreateCollection = "CreateCollection"
	ReadCollection   = "ReadCollection"
	UpdateCollection = "UpdateCollection"
	DeleteCollection = "DeleteCollection"

	AddData    = "addData"
	ReadData   = "readData"
	UpdateData = "updateData"
	DeleteData = "deleteData"
)

const (
	Allow = true
	Deny  = false
)
