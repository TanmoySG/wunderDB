package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type database struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (wh wdbHandlers) CreateDatabase(c *fiber.Ctx) error {
	action := privileges.CreateDatabase

	d := new(database)

	if err := c.BodyParser(d); err != nil {
		return err
	}

	err := wh.wdbClient.AddDatabase(model.Identifier(d.Name))
	resp := response.Format(action, err, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) FetchDatabase(c *fiber.Ctx) error {
	action := privileges.ReadDatabase

	databaseName := c.Params("database")

	fetchedDatabase, err := wh.wdbClient.GetDatabase(model.Identifier(databaseName))
	resp := response.Format(action, err, fetchedDatabase)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) DeleteDatabase(c *fiber.Ctx) error {
	action := privileges.DeleteDatabase

	databaseName := c.Params("database")

	err := wh.wdbClient.DeleteDatabase(model.Identifier(databaseName))
	resp := response.Format(action, err, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
