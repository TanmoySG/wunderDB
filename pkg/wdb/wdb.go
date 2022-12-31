package wdbClient

import (
	d "github.com/TanmoySG/wunderDB/internal/databases"
	r "github.com/TanmoySG/wunderDB/internal/roles"
	u "github.com/TanmoySG/wunderDB/internal/users"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type wdbClient struct {
	HashingAlgorithm string      `json:"hashingAlgorithm"`
	Databases        d.Databases `json:"databases"`
	Roles            r.Roles     `json:"roles"`
	Users            u.Users     `json:"users"`
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
	CreateRole(roleID model.Identifier, allowed []string, denied []string) *er.WdbError
	ListRole() r.Roles
	GrantRoles(userID model.Identifier, permissions []model.Permissions) *er.WdbError
}

func NewWdbClient(databases d.Databases, roles r.Roles, users u.Users, hashingAlgorithm string) Client {
	return wdbClient{
		HashingAlgorithm: hashingAlgorithm,
		Databases:        databases,
		Roles:            roles,
		Users:            users,
	}
}
