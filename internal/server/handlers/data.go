package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

const (
	emptyFilter = ""
)

func (wh wdbHandlers) AddData(c *fiber.Ctx) error {
	privilege := privileges.AddData

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
		incomingData := new(interface{})
		if err := c.BodyParser(incomingData); err != nil {
			return err
		}

		apiError = wh.wdbClient.AddData(model.Identifier(databaseName), model.Identifier(collectionName), incomingData)
	}

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
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

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)

	return nil
}

func (wh wdbHandlers) DeleteData(c *fiber.Ctx) error {
	privilege := privileges.DeleteData

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

		apiError = wh.wdbClient.DeleteData(model.Identifier(databaseName), model.Identifier(collectionName), filter)
	}

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}

func (wh wdbHandlers) UpdateData(c *fiber.Ctx) error {
	privilege := privileges.UpdateData

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
		incomingUpdatedData := new(interface{})
		if err := c.BodyParser(incomingUpdatedData); err != nil {
			return err
		}

		filterKey, filterValue := c.Query("key"), c.Query("value")
		if filterKey == emptyFilter || filterValue == emptyFilter {
			filter = nil
		} else {
			filter = map[string]interface{}{
				"key":   filterKey,
				"value": filterValue,
			}
		}

		apiError = wh.wdbClient.UpdateData(model.Identifier(databaseName), model.Identifier(collectionName), incomingUpdatedData, filter)
	}

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}
