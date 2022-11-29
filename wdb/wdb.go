package wdbClient

import (
	d "github.com/TanmoySG/wunderDB/internal/databases"
	"github.com/TanmoySG/wunderDB/model"
)

type wdbClient struct {
	Databases d.Databases `json:"databases"`
}

type Client interface {
	AddDatabase(databaseId string, metadata model.Metadata) error
	GetDatabase(databaseId string) (*model.Database, error)
	DeleteDatabase(databaseId string) error
}

func NewWdbClient(databases d.Databases) Client {
	return wdbClient{
		Databases: databases,
	}
}
