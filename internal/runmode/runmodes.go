package runmode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
)

type RUN_MODE_TYPE string

type RunModeOpts struct {
	Users  map[model.Identifier]*model.User
	Config config.Config
	app    *fiber.App
}

const (
	RUN_MODE_MAINTENANCE RUN_MODE_TYPE = "RUN_MODE_MAINTENANCE"
	RUN_MODE_NORMAL      RUN_MODE_TYPE = "RUN_MODE_NORMAL"
	RUN_MODE_UPGRADE     RUN_MODE_TYPE = "RUN_MODE_UPGRADE"
)

type INSTRUCTION_TYPE string

const (
	INSTRUCTION_TYPE_CMD INSTRUCTION_TYPE = "INSTRUCTION_TYPE_CMD"
)

type RunInstructions struct {
	ToolName        *string `json:"tool_name"`
	ToolPath        *string `json:"tool_path"`
	ToolArgs        *string `json:"tool_args"`
	InstructionType *string `json:"instruction_type"`
	Instructions    *string `json:"instructions"`
}

// convert to interface and client
func ShouldEnterRunMode(rm RUN_MODE_TYPE) bool {
	switch rm {
	case RUN_MODE_MAINTENANCE:
		return true
	case RUN_MODE_UPGRADE:
		return true
	case RUN_MODE_NORMAL:
		return false
	default:
		return false
	}
}

func (rmo *RunModeOpts) handleAuth(c *fiber.Ctx) (*PatchApiResponse, bool) {
	patchApiResponse := PatchApiResponse{}

	username, password, err := authentication.HandleUserCredentials(c.Get(Authorization))
	if err != nil {
		patchApiResponse.Status = "unauthorized"
		patchApiResponse.Message = "error parsing credentials"
		patchApiResponse.Error = &err.ErrMessage
		patchApiResponse.CommandOutputs = []InstructionOutput{}
		return &patchApiResponse, true
	}

	isAuthorized, err := rmo.AuthenticateUser(model.Identifier(*username), *password)
	if err != nil {
		patchApiResponse.Status = "unauthorized"
		patchApiResponse.Message = "error checking credentials"
		patchApiResponse.Error = &err.ErrMessage
		patchApiResponse.CommandOutputs = []InstructionOutput{}
		return &patchApiResponse, true
	}

	if !isAuthorized {
		patchApiResponse.Status = "unauthorized"
		patchApiResponse.Message = unauthorizedMessage
		patchApiResponse.Error = &unauthorizedMessage
		patchApiResponse.CommandOutputs = []InstructionOutput{}
		return &patchApiResponse, true
	}

	return nil, false
}

func LoadUsers(c config.Config) (map[model.Identifier]*model.User, error) {
	fs := wfs.NewWFileSystem(c.PersistantStoragePath)

	loadedUsers, err := fs.LoadUsers()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	return loadedUsers, nil
}

func (rmo *RunModeOpts) AuthenticateUser(requesterId model.Identifier, password string) (bool, *wdbErrors.WdbError) {
	if requesterId != model.Identifier(rmo.Config.AdminID) {
		return false, &wdbErrors.WdbError{
			ErrCode:        "unauthorizedInMaintenanceMode",
			ErrMessage:     "only admin can authenticate in maintenance mode",
			HttpStatusCode: fiber.StatusUnauthorized,
		}
	}

	user, exists := rmo.Users[requesterId]
	if !exists {
		return authentication.InvalidUser, &wdbErrors.AuthenticatingUserDoesNotExist
	}

	hashedPassword := authentication.Hash(password, user.Authentication.HashingAlgorithm)
	if user.Authentication.HashedSecret == hashedPassword {
		return authentication.ValidUser, nil
	}

	return authentication.InvalidUser, &wdbErrors.InvalidCredentialsError
}

func getErrorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func (rmo *RunModeOpts) sendResponse(c *fiber.Ctx, apiResponse PatchApiResponse, shutdown bool) error {
	c.Set(ContentType, ApplicationJson)

	bytesApiResponse, err := json.Marshal(apiResponse)
	if err != nil {
		return err
	}

	err = c.Send(bytesApiResponse)
	if err != nil {
		return err
	}

	status := fiber.StatusOK
	switch apiResponse.Status {
	case "unauthorized":
		status = fiber.StatusUnauthorized
	case "success":
		status = fiber.StatusOK
	default:
		status = fiber.StatusInternalServerError
	}

	err = c.SendStatus(status)
	if err != nil {
		return err
	}

	// Shutdown the app after the response is sent
	if shutdown {
		go func() {
			fmt.Println("shutting down maintenance mode")
			time.Sleep(1 * time.Second) // Give some time for the response to be sent
			rmo.app.Shutdown()
		}()
	}

	return nil
}
