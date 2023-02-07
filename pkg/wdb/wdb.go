package wdbClient

import (
	"github.com/TanmoySG/wunderDB/internal/config"
	d "github.com/TanmoySG/wunderDB/internal/databases"
	r "github.com/TanmoySG/wunderDB/internal/roles"
	u "github.com/TanmoySG/wunderDB/internal/users"

	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
)

type wdbClient struct {
	HashingAlgorithm string
	Databases        d.Databases
	Roles            r.Roles
	Users            u.Users
	Configurations   model.Configurations
}

type Client interface {
	// Database Methods
	AddDatabase(databaseId model.Identifier) *er.WdbError
	GetDatabase(databaseId model.Identifier) (*model.Database, *er.WdbError)
	DeleteDatabase(databaseId model.Identifier) *er.WdbError

	// Collection Methods
	AddCollection(databaseId, collectionId model.Identifier, schema model.Schema) *er.WdbError
	GetCollection(databaseId, collectionId model.Identifier) (*model.Collection, *er.WdbError)
	DeleteCollection(databaseId, collectionId model.Identifier) *er.WdbError

	// Data Methods
	AddData(databaseId, collectionId model.Identifier, inputData interface{}) *er.WdbError
	GetData(databaseId model.Identifier, collectionId model.Identifier, filters interface{}) (map[model.Identifier]*model.Datum, *er.WdbError)
	UpdateData(databaseId model.Identifier, collectionId model.Identifier, updatedData interface{}, filters interface{}) *er.WdbError
	DeleteData(databaseId model.Identifier, collectionId model.Identifier, filters interface{}) *er.WdbError

	// Methods for Roles and Users
	CreateUser(userID model.Identifier, password string) *er.WdbError
	AuthenticateUser(userID model.Identifier, password string) (bool, *er.WdbError)
	CreateRole(roleID model.Identifier, allowed []string, denied []string) *er.WdbError
	ListRole() r.Roles
	CheckUserPermissions(userID model.Identifier, privilege string, entities model.Entities) (bool, *er.WdbError)
	GrantRoles(userID model.Identifier, permissions model.Permissions) *er.WdbError

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
	}
}
