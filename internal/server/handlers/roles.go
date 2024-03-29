package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

type role struct {
	Role    string   `json:"role" xml:"role" form:"role" validate:"required"`
	Allowed []string `json:"allowed" xml:"allowed" form:"allowed" validate:"required"`
	Denied  []string `json:"denied" xml:"denied" form:"denied"`
	Hidden  bool     `json:"hidden" xml:"hidden" form:"hidden"`
}

func (wh wdbHandlers) CreateRole(c *fiber.Ctx) error {
	privilege := privileges.CreateRole

	var apiError *er.WdbError

	r := new(role)
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := validateRequest(r); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
		if !isValid {
			apiError = error
		} else {
			apiError = wh.wdbClient.CreateRole(model.Identifier(r.Role), r.Allowed, r.Denied, r.Hidden)
		}
	}

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}

func (wh wdbHandlers) UpdateRole(c *fiber.Ctx) error {
	privilege := privileges.UpdateRole

	var apiError *er.WdbError

	r := new(role)
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := validateRequest(r); err != nil {
		apiError = err
	} else {
		isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
		if !isValid {
			apiError = error
		} else {
			apiError = wh.wdbClient.UpdateRole(model.Identifier(r.Role), r.Allowed, r.Denied, r.Hidden)
		}
	}

	resp := response.Format(privilege, apiError, nil, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}

func (wh wdbHandlers) ListRoles(c *fiber.Ctx) error {
	privilege := privileges.ListRole

	forceListFlag := c.Query("force", "false")

	var apiError *er.WdbError
	var roleList roles.Roles

	isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
	if !isValid {
		apiError = error
	} else {
		actorUserId := authentication.GetActor(c.Get(Authorization))
		roleList, apiError = wh.wdbClient.ListRole(actorUserId, forceListFlag)
	}

	resp := response.Format(privilege, apiError, roleList, *wh.notices...)

	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}
