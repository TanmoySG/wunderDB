package handlers

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type role struct {
	Role    string   `json:"role" xml:"role" form:"role"`
	Allowed []string `json:"allowed" xml:"allowed" form:"allowed"`
	Denied  []string `json:"denied" xml:"denied" form:"denied"`
}

func (wh wdbHandlers) CreateRole(c *fiber.Ctx) error {
	privilege := privileges.CreateRole

	var apiError *er.WdbError

	r := new(role)
	if err := c.BodyParser(r); err != nil {
		return err
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
	if !isValid {
		apiError = error
	} else {
		apiError = wh.wdbClient.CreateRole(model.Identifier(r.Role), r.Allowed, r.Denied)
	}

	resp := response.Format(privilege, apiError, nil)

	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}

func (wh wdbHandlers) ListRoles(c *fiber.Ctx) error {
	privilege := privileges.ListRole

	var apiError *er.WdbError
	var roleList roles.Roles

	isValid, error := wh.handleAuthenticationAndAuthorization(c, noEntities, privilege)
	if !isValid {
		apiError = error
	} else {
		roleList = wh.wdbClient.ListRole()
	}

	resp := response.Format(privilege, apiError, roleList)

	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}
