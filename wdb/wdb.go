package wdbClient

import "github.com/TanmoySG/wunderDB/model"

type client struct {
	Databases map[model.Identifier]*model.Database
}

type Client interface {
	AddDatabase()
}

func NewWDBClient(databases map[model.Identifier]*model.Database) Client {
	return client{
		Databases: databases,
	}
}
