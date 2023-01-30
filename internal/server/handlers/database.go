package handlers

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type database struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}

func (wh wdbHandlers) CreateDatabase(c *fiber.Ctx) error {
	privilege := privileges.CreateDatabase

	var apiError *er.WdbError

	d := new(database)
	if err := c.BodyParser(d); err != nil {
		return err
	}

	if err := ValidateRequest(d); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
		if !isValid {
			apiError = error
		} else {
			apiError = wh.wdbClient.AddDatabase(model.Identifier(d.Name))
		}
	}
	resp := response.Format(privilege, apiError, nil)

	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}

func (wh wdbHandlers) FetchDatabase(c *fiber.Ctx) error {
	privilege := privileges.ReadDatabase

	var apiError *er.WdbError
	var fetchedDatabase *model.Database

	databaseName := c.Params("database")
	entities := model.Entities{
		Databases: &databaseName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		fetchedDatabase, apiError = wh.wdbClient.GetDatabase(model.Identifier(databaseName))
	}

	resp := response.Format(privilege, apiError, fetchedDatabase)

	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}

func (wh wdbHandlers) DeleteDatabase(c *fiber.Ctx) error {
	privilege := privileges.DeleteDatabase

	var apiError *er.WdbError

	databaseName := c.Params("database")
	entities := model.Entities{
		Databases: &databaseName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		apiError = wh.wdbClient.DeleteDatabase(model.Identifier(databaseName))
	}

	resp := response.Format(privilege, apiError, nil)

	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}
