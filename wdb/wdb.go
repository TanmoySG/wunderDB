package wdbClient

import (
	d "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/model"
)

type wdbClient struct {
	Databases d.Databases `json:"databases"`
}

type Client interface {
	// Database Methods
	AddDatabase(databaseId model.Identifier, metadata model.Metadata) error
	GetDatabase(databaseId model.Identifier) (*model.Database, error)
	DeleteDatabase(databaseId model.Identifier) error

	// Collection Methods
	AddCollection(databaseId, collectionId model.Identifier, schema model.Schema, metadata model.Metadata) error
}

func NewWdbClient(databases d.Databases) Client {
	return wdbClient{
		Databases: databases,
	}
}
