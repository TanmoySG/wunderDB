package handlers

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type collection struct {
	Name   string                 `json:"name" xml:"name" form:"name"`
	Schema map[string]interface{} `json:"schema" xml:"schema" form:"schema"`
}

func (wh wdbHandlers) CreateCollection(c *fiber.Ctx) error {
	privilege := privileges.CreateCollection

	var apiError *er.WdbError

	databaseName := c.Params("database")

	collection := new(collection)
	if err := c.BodyParser(collection); err != nil {
		return err
	}

	entities := model.Entities{
		Databases: &databaseName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		apiError = wh.wdbClient.AddCollection(model.Identifier(databaseName), model.Identifier(collection.Name), collection.Schema)
	}

	resp := response.Format(privilege, apiError, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) FetchCollection(c *fiber.Ctx) error {
	privilege := privileges.ReadCollection

	var apiError *er.WdbError
	var fetchedDatabase *model.Collection

	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	entities := model.Entities{
		Databases:   &databaseName,
		Collections: &collectionName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		fetchedDatabase, apiError = wh.wdbClient.GetCollection(model.Identifier(databaseName), model.Identifier(collectionName))
	}

	resp := response.Format(privilege, apiError, fetchedDatabase)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) DeleteCollection(c *fiber.Ctx) error {
	privilege := privileges.DeleteCollection

	var apiError *er.WdbError

	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	entities := model.Entities{
		Databases:   &databaseName,
		Collections: &collectionName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		apiError = wh.wdbClient.DeleteCollection(model.Identifier(databaseName), model.Identifier(collectionName))
	}

	resp := response.Format(privilege, apiError, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
