package routes

const (
	CreateDatabase = "/databases"
	FetchDatabase  = "/databases/:database"
	DeleteDatabase = "/databases/:database"

	// Collection Routes
	CreateCollection = "/databases/:database/collections"
	FetchCollection  = "/databases/:database/collections/:collection"
	DeleteCollection = "/databases/:database/collections/:collection"

	// Records Routes
	AddRecords    = "/databases/:database/collections/:collection/records"
	ReadRecords   = "/databases/:database/collections/:collection/records"
	QueryRecords  = "/databases/:database/collections/:collection/records/query" // route for executing jsonpath queries
	DeleteRecords = "/databases/:database/collections/:collection/records"
	UpdateRecords = "/databases/:database/collections/:collection/records"

	// Record Routes
	ReadRecord   = "/databases/:database/collections/:collection/records/:id"
	DeleteRecord = "/databases/:database/collections/:collection/records/:id"
	UpdateRecord = "/databases/:database/collections/:collection/records/:id"

	// Role Routes
	CreateRole = "/roles"
	ListRoles  = "/roles"
	UpdateRole = "/roles"

	// User Routes
	CreateUser = "/users"
	LoginUser  = "/users/login"
	GrantRole  = "/users/grant"
	RevokeRole = "/users/revoke"
)
