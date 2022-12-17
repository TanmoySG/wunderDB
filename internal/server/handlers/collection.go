package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type collection struct {
	Name   string                 `json:"name" xml:"name" form:"name"`
	Schema map[string]interface{} `json:"schema" xml:"schema" form:"schema"`
}

func (wh wdbHandlers) CreateCollection(c *fiber.Ctx) error {
	databaseName := c.Params("database")

	collection := new(collection)
	if err := c.BodyParser(collection); err != nil {
		return err
	}

	err := wh.wdbClient.AddCollection(model.Identifier(databaseName), model.Identifier(collection.Name), collection.Schema)
	resp := response.Format(CreateCollectionAction, err, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) FetchCollection(c *fiber.Ctx) error {
	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	fetchedDatabase, err := wh.wdbClient.GetCollection(model.Identifier(databaseName), model.Identifier(collectionName))
	resp := response.Format(FetchCollectionAction, err, fetchedDatabase)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) DeleteCollection(c *fiber.Ctx) error {
	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	err := wh.wdbClient.DeleteCollection(model.Identifier(databaseName), model.Identifier(collectionName))
	resp := response.Format(DeleteCollectionAction, err, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
