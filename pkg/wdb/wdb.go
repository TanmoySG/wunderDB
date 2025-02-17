package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	d "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/internal/records"
	r "github.com/TanmoySG/wunderDB/internal/roles"
	u "github.com/TanmoySG/wunderDB/internal/users"
	"github.com/TanmoySG/wunderDB/model/redacted"

	"github.com/TanmoySG/wunderDB/model"
	"github.com/TanmoySG/wunderDB/pkg/utils/safename"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

const (
	// only alphanumeric, hyphen, underscore and period allowed
	safeNamePattern = "(?misU)^[a-zA-Z0-9-_.@]+$"
)

type wdbClient struct {
	HashingAlgorithm string
	Databases        d.Databases
	Roles            r.Roles
	Users            u.Users
	Configurations   model.Configurations
	safeName         safename.SafeNameClient
}

type Client interface {
	// Database Methods
	AddDatabase(databaseId model.Identifier, userId model.Identifier) *er.WdbError
	GetDatabase(databaseId model.Identifier) (*redacted.RedactedD, *er.WdbError)
	DeleteDatabase(databaseId model.Identifier) *er.WdbError

	// Collection Methods
	AddCollection(databaseId, collectionId model.Identifier, schema model.Schema, primaryKey *model.Identifier) *er.WdbError
	GetCollection(databaseId, collectionId model.Identifier) (*redacted.RedactedC, *er.WdbError)
	DeleteCollection(databaseId, collectionId model.Identifier) *er.WdbError

	// Data Methods
	AddRecords(databaseId, collectionId model.Identifier, inputData interface{}) *er.WdbError
	GetRecords(databaseId model.Identifier, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Record, *er.WdbError)
	UpdateRecords(databaseId model.Identifier, collectionId model.Identifier, updatedData interface{}, filters interface{}) *er.WdbError
	DeleteRecords(databaseId model.Identifier, collectionId model.Identifier, filters interface{}) *er.WdbError
	QueryRecords(databaseId, collectionId model.Identifier, query string, mode records.QueryType) (interface{}, *er.WdbError)

	// Users Methods
	CreateUser(userID model.Identifier, password string, metadata model.Metadata) *er.WdbError
	AuthenticateUser(userID model.Identifier, password string) (bool, *er.WdbError)
	CheckUserPermissions(userID model.Identifier, privilege string, entities model.Entities) (bool, *er.WdbError)
	GrantRole(userID model.Identifier, permissions model.Permissions) *er.WdbError
	RevokeRole(userID model.Identifier, permission model.Permissions) (map[string]int, *er.WdbError)

	// Roles Methods
	ListRole(requesterId, forceListAllRoles string) (r.Roles, *er.WdbError)
	CreateRole(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError
	UpdateRole(roleID model.Identifier, allowed []string, denied []string, hidden bool) *er.WdbError

	// Admin Method
	InitializeAdmin(config *config.Config)
}

func NewWdbClient(configurations model.Configurations, databases d.Databases, roles r.Roles, users u.Users, hashingAlgorithm string) Client {
	return wdbClient{
		HashingAlgorithm: hashingAlgorithm,
		Databases:        databases,
		Roles:            roles,
		Users:            users,
		Configurations:   configurations,
		// safeNamePattern will always compile safely. Replace
		// with UsePattern when NewWdbClient returns error too.
		safeName: safename.MustUsePattern(safeNamePattern),
	}
}
