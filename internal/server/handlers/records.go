package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/records"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

const (
	emptyFilter = ""
	idKey       = "id"
)

type queryRequest struct {
	QueryMode   string `json:"mode" xml:"mode" form:"mode" validate:"required"`
	QueryString string `json:"query" xml:"query" form:"query" validate:"required"`
}

func (wh wdbHandlers) AddRecords(c *fiber.Ctx) error {
	privilege := privileges.AddRecords

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

		apiError = wh.wdbClient.AddRecords(model.Identifier(databaseName), model.Identifier(collectionName), incomingData)
	}

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}

func (wh wdbHandlers) ReadRecords(c *fiber.Ctx) error {
	privilege := privileges.ReadRecords

	var apiError *er.WdbError
	var fetchedData map[model.Identifier]*model.Record
	var filter interface{}

	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	id := c.Params("id")

	entities := model.Entities{
		Databases:   &databaseName,
		Collections: &collectionName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		filterKey, filterValue := c.Query("key"), c.Query("value")
		if id != emptyFilter {
			filter = map[string]interface{}{
				"key":   idKey,
				"value": id,
			}
		} else if filterKey != emptyFilter || filterValue != emptyFilter {
			filter = map[string]interface{}{
				"key":   filterKey,
				"value": filterValue,
			}
		} else {
			filter = nil
		}

		fetchedData, apiError = wh.wdbClient.GetRecords(model.Identifier(databaseName), model.Identifier(collectionName), filter)
	}

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)

	return nil
}

func (wh wdbHandlers) QueryRecords(c *fiber.Ctx) error {
	privilege := privileges.QueryRecords

	var apiError *er.WdbError
	var queriedData interface{}
	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	entities := model.Entities{
		Databases:   &databaseName,
		Collections: &collectionName,
	}

	query := new(queryRequest)
	if err := c.BodyParser(query); err != nil {
		return err
	}

	if err := validateRequest(query); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
		if !isValid {
			apiError = error
		} else {
			queriedData, apiError = wh.wdbClient.QueryRecords(
				model.Identifier(databaseName),
				model.Identifier(collectionName),
				query.QueryString,
				records.QueryType(query.QueryMode),
			)
		}
	}

	resp := response.Format(privilege, apiError, queriedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)

	return nil
}

func (wh wdbHandlers) DeleteRecords(c *fiber.Ctx) error {
	privilege := privileges.DeleteRecords

	var apiError *er.WdbError
	var fetchedData map[model.Identifier]*model.Record
	var filter interface{}

	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	id := c.Params("id")

	entities := model.Entities{
		Databases:   &databaseName,
		Collections: &collectionName,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		filterKey, filterValue := c.Query("key"), c.Query("value")
		if id != emptyFilter {
			filter = map[string]interface{}{
				"key":   idKey,
				"value": id,
			}
		} else if filterKey != emptyFilter || filterValue != emptyFilter {
			filter = map[string]interface{}{
				"key":   filterKey,
				"value": filterValue,
			}
		} else {
			filter = nil
		}

		apiError = wh.wdbClient.DeleteRecords(model.Identifier(databaseName), model.Identifier(collectionName), filter)
	}

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}

func (wh wdbHandlers) UpdateRecords(c *fiber.Ctx) error {
	privilege := privileges.UpdateRecords

	var apiError *er.WdbError
	var fetchedData map[model.Identifier]*model.Record
	var filter interface{}

	databaseName := c.Params("database")
	collectionName := c.Params("collection")

	id := c.Params("id")

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
		if id != emptyFilter {
			filter = map[string]interface{}{
				"key":   idKey,
				"value": id,
			}
		} else if filterKey != emptyFilter || filterValue != emptyFilter {
			filter = map[string]interface{}{
				"key":   filterKey,
				"value": filterValue,
			}
		} else {
			filter = nil
		}

		apiError = wh.wdbClient.UpdateRecords(model.Identifier(databaseName), model.Identifier(collectionName), incomingUpdatedData, filter)
	}

	resp := response.Format(privilege, apiError, fetchedData, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, entities)
	return nil
}
