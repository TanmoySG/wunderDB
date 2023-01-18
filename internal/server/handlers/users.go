package handlers

import (
	"strconv"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type userPermissions struct {
	Username   string            `json:"username" xml:"username" form:"username"`
	Permission model.Permissions `json:"permissions" xml:"permissions" form:"permissions"`
}

type newUser struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	// HashingAlgorithm string `json:"hashingAlgorithm" xml:"hashingAlgorithm" form:"hashingAlgorithm"`
}

func (wh wdbHandlers) CreateUser(c *fiber.Ctx) error {
	privilege := privileges.CreateUser

	u := new(newUser)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	error := wh.wdbClient.CreateUser(model.Identifier(u.Username), u.Password)
	resp := response.Format(privilege, error, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) GrantRoles(c *fiber.Ctx) error {
	privilege := privileges.GrantRole

	var data map[string]interface{}
	var apiError *er.WdbError

	u := new(userPermissions)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	entities := model.Entities{
		Databases:   u.Permission.On.Databases,
		Collections: u.Permission.On.Collections,
	}

	isValid, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	if !isValid {
		apiError = error
	} else {
		apiError = wh.wdbClient.GrantRoles(model.Identifier(u.Username), u.Permission)
	}

	resp := response.Format(privilege, apiError, data)
	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) CheckPermissions(c *fiber.Ctx) error {
	privilege := c.Query("privilege")
	database := c.Query("database")
	collection := c.Query("collection")

	entities := model.Entities{
		Databases:   &database,
		Collections: &collection,
	}

	authStatus, error := wh.handleAuthenticationAndAuthorization(c, entities, privilege)
	data := map[string]string{
		"privilege": privilege,
		"allowed":   strconv.FormatBool(authStatus),
	}
	resp := response.Format("", error, data)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
