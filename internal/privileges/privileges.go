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

func IsAvailable(privilege string) bool {

	availableActions := []string{
		Wildcard,
		CreateRole,
		GrantRole,
		UpdateRole,
		CreateDatabase,
		ReadDatabase,
		UpdateDatabase,
		DeleteDatabase,
		CreateCollection,
		ReadCollection,
		UpdateCollection,
		DeleteCollection,
		AddData,
		ReadData,
		UpdateData,
		DeleteData,
	}

	for _, action := range availableActions {
		if privilege == action {
			return true
		}
	}

	return false
}
