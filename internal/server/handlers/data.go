package handlers

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

const (
	emptyFilter = ""
)

func (wh wdbHandlers) AddData(c *fiber.Ctx) error {
	privilege := privileges.AddData

	var apiError *er.WdbError

	incomingData := new(interface{})
	if err := c.BodyParser(incomingData); err != nil {
		return err
	}

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
		apiError = wh.wdbClient.AddData(model.Identifier(databaseName), model.Identifier(collectionName), incomingData)
	}

	resp := response.Format(privilege, apiError, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) ReadData(c *fiber.Ctx) error {
	privilege := privileges.ReadData

	var apiError *er.WdbError
	var fetchedData map[model.Identifier]*model.Datum
	var filter interface{}

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
		filterKey, filterValue := c.Query("key"), c.Query("value")
		if filterKey == emptyFilter || filterValue == emptyFilter {
			filter = nil
		} else {
			filter = map[string]interface{}{
				"key":   filterKey,
				"value": filterValue,
			}
		}

		fetchedData, apiError = wh.wdbClient.GetData(model.Identifier(databaseName), model.Identifier(collectionName), filter)
	}

	resp := response.Format(privilege, apiError, fetchedData)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

// func (wh wdbHandlers) DeleteData(c *fiber.Ctx) error {
// 	privilege := privileges.DeleteDatabase

// 	var apiError *er.WdbError

// 	databaseName := c.Params("database")
// 	entities := model.Entities{
// 		Databases: &databaseName,
// 	}

// 	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
// 	if !isValid {
// 		apiError = error
// 	} else {
// 		apiError = wh.wdbClient.DeleteDatabase(model.Identifier(databaseName))
// 	}

// 	resp := response.Format(privilege, apiError, nil)

// 	c.Send(resp.Marshal())
// 	return c.SendStatus(resp.HttpStatusCode)
// }

// func (wh wdbHandlers) UpdateData(c *fiber.Ctx) error {
// 	privilege := privileges.DeleteDatabase

// 	var apiError *er.WdbError

// 	databaseName := c.Params("database")
// 	entities := model.Entities{
// 		Databases: &databaseName,
// 	}

// 	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
// 	if !isValid {
// 		apiError = error
// 	} else {
// 		apiError = wh.wdbClient.DeleteDatabase(model.Identifier(databaseName))
// 	}

// 	resp := response.Format(privilege, apiError, nil)

// 	c.Send(resp.Marshal())
// 	return c.SendStatus(resp.HttpStatusCode)
// }
