package wdbClient

import (
	d "github.com/TanmoySG/wunderDB/internal/databases"
	r "github.com/TanmoySG/wunderDB/internal/roles"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/model"
)

type wdbClient struct {
	Databases d.Databases `json:"databases"`
	Roles     r.Roles     `json:"roles"`
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
	CreateRole(roleID model.Identifier, allowedActions []string, deniedActions []string) *er.WdbError
	ListRole() r.Roles
}

func NewWdbClient(databases d.Databases, roles r.Roles) Client {
	return wdbClient{
		Databases: databases,
		Roles:     roles,
	}
}
