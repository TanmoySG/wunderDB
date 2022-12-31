package handlers

import (
	"strconv"

	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Username    string              `json:"username" xml:"username" form:"username"`
	Permissions []model.Permissions `json:"permissions" xml:"permissions" form:"permissions"`
}

func (wh wdbHandlers) CreateUser(c *fiber.Ctx) error {
	action := privileges.CreateUser

	username, password, _ := authentication.HandleUserCredentials(c.Get(Authorization))

	error := wh.wdbClient.CreateUser(model.Identifier(*username), *password)
	resp := response.Format(action, error, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) GrantRoles(c *fiber.Ctx) error {
	action := privileges.GrantRole

	_, _, err := authentication.HandleUserCredentials(c.Get(Authorization))
	if err != nil {
		return err
	}

	u := new(user)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	error := wh.wdbClient.GrantRoles(model.Identifier(u.Username), u.Permissions)
	resp := response.Format(action, error, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}

func (wh wdbHandlers) CheckPermissions(c *fiber.Ctx) error {
	privilege := c.Query("privilege")
	database := c.Query("database")
	collection := c.Query("collection")

	username, _, err := authentication.HandleUserCredentials(c.Get(Authorization))
	if err != nil {
		return err
	}

	entities := model.Entities{
		Databases:   &database,
		Collections: &collection,
	}

	isAllowed, error := wh.wdbClient.CheckUserPermissions(model.Identifier(*username), privilege, entities)
	data := map[string]string{
		"privilege": privilege,
		"allowed":   strconv.FormatBool(*isAllowed),
	}
	resp := response.Format("", error, data)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
