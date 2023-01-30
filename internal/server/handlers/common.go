package handlers

import (
	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	noEntities = model.Entities{}
)

const (
	authSuccessful = true
	authFailure    = false
)

func (wh wdbHandlers) handleAuthenticationAndAuthorization(c *fiber.Ctx, entities model.Entities, privilege string) (bool, *er.WdbError) {

	username, isValidUser, error := wh.handleAuthentication(c)
	if !isValidUser {
		return authFailure, error
	}

	isValid, error := wh.handleAuthorization(*username, entities, privilege)

	if !isValid {
		return authFailure, error
	}

	return authSuccessful, nil
}

func (wh wdbHandlers) handleAuthentication(c *fiber.Ctx) (*string, bool, *er.WdbError) {
	username, password, error := authentication.HandleUserCredentials(c.Get(Authorization))
	if error != nil {
		return nil, authFailure, error
	}

	isValid, error := wh.wdbClient.AuthenticateUser(model.Identifier(*username), *password)
	if error != nil {
		return nil, authFailure, error
	}

	if !isValid {
		return nil, authFailure, error
	}
	return username, authSuccessful, nil
}

func (wh wdbHandlers) handleAuthorization(username string, entity model.Entities, privilege string) (bool, *er.WdbError) {

	isAllowed, error := wh.wdbClient.CheckUserPermissions(model.Identifier(username), privilege, entity)

	if error != nil {
		return authFailure, error
	}

	if !isAllowed {
		return authFailure, error
	}

	return authSuccessful, nil
}

func SendResponse(c *fiber.Ctx, marshaledResponse []byte, statusCode int) error {
	c.Set(ContentType, ApplicationJson)
	c.Send(marshaledResponse)
	return c.SendStatus(statusCode)
}

func ValidateRequest(request any) *er.WdbError {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		return &er.ValidationError
	}
	return nil
}
