package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

type collection struct {
	Name   string                 `json:"name" xml:"name" form:"name" validate:"required"`
	Schema map[string]interface{} `json:"schema" xml:"schema" form:"schema" validate:"required"`
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
		Databases:   &databaseName,
		Collections: &collection.Name, // check
	}

	if err := validateRequest(collection); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
		if !isValid {
			apiError = error
		} else {
			apiError = wh.wdbClient.AddCollection(model.Identifier(databaseName), model.Identifier(collection.Name), collection.Schema)
		}
	}

	resp := response.Format(privilege, apiError, nil)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
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

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
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

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}
