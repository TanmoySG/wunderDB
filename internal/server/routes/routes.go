package routes

const (
	CreateDatabaseRoute = "/api/databases"
	FetchDatabaseRoute  = "/api/databases/:database"
	DeleteDatabase      = "/api/databases/:database"

	// Collection Routes
	CreateCollection = "/api/databases/:database/collections"
	FetchCollection  = "/api/databases/:database/collections/:collection"
	DeleteCollection = "/api/databases/:database/collections/:collection"

	// Data Routes
	AddData    = "/api/databases/:database/collections/:collection/data"
	ReadData   = "/api/databases/:database/collections/:collection/data"
	DeleteData = "/api/databases/:database/collections/:collection/data"
	UpdateData = "/api/databases/:database/collections/:collection/data"

	// Role Routes
	CreateRole = "/api/roles"
	ListRoles  = "/api/roles"

	// User Routes
	CreateUser = "/api/users"
	GrantRoles = "/api/users/grant"
)
