package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

type database struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required"`
}

func (wh wdbHandlers) CreateDatabase(c *fiber.Ctx) error {
	privilege := privileges.CreateDatabase
	var apiError *er.WdbError

	database := new(database)
	if err := c.BodyParser(database); err != nil {
		return err
	}

	entities := model.Entities{
		Databases: &database.Name,
	}

	if err := validateRequest(database); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
		if !isValid {
			apiError = error
		} else {
			actorUserId := authentication.GetActor(c.Get(Authorization))
			apiError = wh.wdbClient.AddDatabase(model.Identifier(database.Name), model.Identifier(actorUserId))
		}
	}
	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)

	return nil
}

func (wh wdbHandlers) FetchDatabase(c *fiber.Ctx) error {
	privilege := privileges.ReadDatabase

	var apiError *er.WdbError
	var fetchedDatabase *redacted.RedactedD

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

	resp := response.Format(privilege, apiError, fetchedDatabase, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)

	return nil
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

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}
