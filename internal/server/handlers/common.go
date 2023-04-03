package handlers

import (
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/internal/txlogs"
	txlModel "github.com/TanmoySG/wunderDB/internal/txlogs/model"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/model"
	er "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	AuthorizationHeader = "Authorization"

	authSuccessful = true
	authFailure    = false
)

var (
	noEntities = model.Entities{}
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

func SendResponse(c *fiber.Ctx, apiResponse response.ApiResponse) error {
	c.Set(ContentType, ApplicationJson)

	marshaledResponse := apiResponse.Marshal()
	err := c.Send(marshaledResponse)
	if err != nil {
		return err
	}

	err = c.SendStatus(apiResponse.HttpStatusCode)
	if err != nil {
		return err
	}

	return nil
}

func ValidateRequest(request any) *er.WdbError {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		return &er.ValidationError
	}
	return nil
}

func (wh wdbHandlers) handleTransactions(c *fiber.Ctx, apiResponse response.ApiResponse, entities model.Entities) error {
	if txlogs.IsTxnLoggable(apiResponse.Response.Action) {
		if apiResponse.Response.Status == response.StatusSuccess {
			databaseEntity := *entities.Databases
			if entities.Databases == nil {
				databaseEntity = ""
			}

			txnActor := txlogs.GetTxnActor(c.Get(AuthorizationHeader))
			txnAction := apiResponse.Response.Action

			txnHttpDetails := txlogs.GetTxnHttpDetails(*c)
			txnEntityPath := txlModel.TxlogSchemaJsonEntityPath{
				Database:   databaseEntity,
				Collection: entities.Collections,
			}

			txnLog := txlogs.CreateTxLog(txnAction, txnActor, apiResponse.Response.Status, txnEntityPath, txnHttpDetails)

			// dotTxLogs := txlogs.UseDotTxLog(".") // move to config/init
			err := wh.wdbTxLogs.Log(txnLog)
			if err != nil {
				return err
			}
			err = wh.wdbTxLogs.Commit()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
