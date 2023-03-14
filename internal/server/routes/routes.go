package routes

const (
	CreateDatabase = "/databases"
	FetchDatabase  = "/databases/:database"
	DeleteDatabase = "/databases/:database"

	// Collection Routes
	CreateCollection = "/databases/:database/collections"
	FetchCollection  = "/databases/:database/collections/:collection"
	DeleteCollection = "/databases/:database/collections/:collection"

	// Data Routes
	AddData    = "/databases/:database/collections/:collection/data"
	ReadData   = "/databases/:database/collections/:collection/data"
	DeleteData = "/databases/:database/collections/:collection/data"
	UpdateData = "/databases/:database/collections/:collection/data"

	// Role Routes
	CreateRole = "/roles"
	ListRoles  = "/roles"

	// User Routes
	CreateUser = "/users"
	LoginUser  = "/users/login"
	GrantRoles = "/users/grant"
)
