package handlers

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

type userPermissions struct {
	Username   string            `json:"username" xml:"username" form:"username" validate:"required"`
	Permission model.Permissions `json:"permissions" xml:"permissions" form:"permissions" validate:"required,dive"`
}

type newUser struct {
	Username string `json:"username" xml:"username" form:"username" validate:"required"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}

func (wh wdbHandlers) LoginUser(c *fiber.Ctx) error {
	privilege := privileges.LoginUser

	data := "user not logged-in"

	username, isValidated, apiError := wh.handleAuthentication(c)
	if isValidated {
		data = fmt.Sprintf("%s logged-in", *username)
	}

	resp := response.Format(privilege, apiError, data, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)

	return nil
}

func (wh wdbHandlers) CreateUser(c *fiber.Ctx) error {
	privilege := privileges.CreateUser
	var apiError *er.WdbError

	u := new(newUser)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	if err := validateRequest(u); err != nil {
		apiError = err
	} else {
		apiError = wh.wdbClient.CreateUser(model.Identifier(u.Username), u.Password)
	}

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}

func (wh wdbHandlers) GrantRole(c *fiber.Ctx) error {
	privilege := privileges.GrantRole

	var entities model.Entities
	var data map[string]interface{}
	var apiError *er.WdbError

	up := new(userPermissions)

	if err := c.BodyParser(up); err != nil {
		return err
	}

	if err := validateRequest(up); err != nil {
		apiError = err
	} else {
		if up.Permission.On != nil {
			entities = model.Entities{
				Databases:   up.Permission.On.Databases,
				Collections: up.Permission.On.Collections,
			}
		}

		isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
		if !isValid {
			apiError = error
		} else {
			apiError = wh.wdbClient.GrantRole(model.Identifier(up.Username), up.Permission)
		}
	}

	resp := response.Format(privilege, apiError, data, *wh.notices...)
	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}

func (wh wdbHandlers) RevokeRole(c *fiber.Ctx) error {
	privilege := privileges.RevokeRole

	var entities model.Entities
	var data map[string]int
	var apiError *er.WdbError

	up := new(userPermissions)

	if err := c.BodyParser(up); err != nil {
		return err
	}

	if err := validateRequest(up); err != nil {
		apiError = err
	} else {
		if up.Permission.On != nil {
			entities = model.Entities{
				Databases:   up.Permission.On.Databases,
				Collections: up.Permission.On.Collections,
			}
		}

		isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
		if !isValid {
			apiError = error
		} else {
			data, apiError = wh.wdbClient.RevokeRole(model.Identifier(up.Username), up.Permission)
		}
	}

	resp := response.Format(privilege, apiError, data, *wh.notices...)
	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}
