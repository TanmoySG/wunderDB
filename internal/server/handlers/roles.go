package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/roles"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type role struct {
	Role        string            `json:"role" xml:"role" form:"role"`
	Permissions roles.Permissions `json:"permissions" xml:"permissions" form:"permissions"`
}

func (wh wdbHandlers) CreateRole(c *fiber.Ctx) error {
	action := privileges.CreateRole

	r := new(role)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	err := wh.wdbClient.CreateRole(model.Identifier(r.Role), r.Permissions)
	resp := response.Format(action, err, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) ListRoles(c *fiber.Ctx) error {
	action := privileges.CreateRole

	r := new(role)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	roleList := wh.wdbClient.ListRole()
	resp := response.Format(action, nil, roleList)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
