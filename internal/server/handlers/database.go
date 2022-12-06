package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type database struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (wh wdbHandlers) CreateDatabase(c *fiber.Ctx) error {
	d := new(database)

	if err := c.BodyParser(d); err != nil {
		return err
	}

	err := wh.wdbClient.AddDatabase(model.Identifier(d.Name))
	resp := response.Format(CreateDatabaseAction, err, nil)

	return c.Send(resp.Marshal())
}

func (wh wdbHandlers) FetchDatabase(c *fiber.Ctx) error {
	databaseName := c.Params("database")

	fetchedDatabase, err := wh.wdbClient.GetDatabase(model.Identifier(databaseName))
	resp := response.Format(FetchDatabaseAction, err, fetchedDatabase)

	return c.Send(resp.Marshal())
}

func (wh wdbHandlers) DeleteDatabase(c *fiber.Ctx) error {
	databaseName := c.Params("database")

	err := wh.wdbClient.DeleteDatabase(model.Identifier(databaseName))
	resp := response.Format(DeleteDatabaseAction, err, nil)

	return c.Send(resp.Marshal())
}
