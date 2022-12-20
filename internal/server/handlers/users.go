package handlers

import (
	"encoding/base64"
	"strings"

	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

func (wh wdbHandlers) CreateUser(c *fiber.Ctx) error {
	action := privileges.CreateUser

	authorizationHeader := strings.Split(c.Get(Authorization), " ")

	decodedCredentials, err := base64.StdEncoding.DecodeString(authorizationHeader[1])
	if err != nil {
		return err
	}

	credentialArray := strings.Split(string(decodedCredentials), ":")

	username, password := credentialArray[0], credentialArray[1]
	r := new(role)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	error := wh.wdbClient.CreateUser(model.Identifier(username), password)
	resp := response.Format(action, error, nil)

	c.Send(resp.Marshal())
	return c.SendStatus(resp.HttpStatusCode)
}
