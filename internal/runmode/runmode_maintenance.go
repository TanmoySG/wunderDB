package runmode

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/server/middlewares/recovery"
	"github.com/TanmoySG/wunderDB/internal/users/authentication"
	"github.com/TanmoySG/wunderDB/internal/wfs"
	"github.com/TanmoySG/wunderDB/model"
	wdbErrors "github.com/TanmoySG/wunderDB/pkg/wdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	Authorization = "Authorization"
	ContentType   = "Content-Type"

	ApplicationJson = "application/json"

	shutdown = true
)

var (
	unauthorizedMessage string = "user is not authorized to perform this action"
)

type Instruction struct {
	Commands []string `json:"commands" xml:"commands" form:"commands" validate:"required"`
}

type InstructionOutput struct {
	Command   string `json:"command" xml:"command" form:"command"`
	Output    string `json:"output" xml:"output" form:"output"`
	Error     string `json:"error" xml:"error" form:"error"`
	IsSuccess bool   `json:"isSuccess" xml:"isSuccess" form:"isSuccess"`
}

type PatchApiResponse struct {
	Message        string              `json:"message" xml:"message" form:"message"`
	Status         string              `json:"status" xml:"status" form:"status"`
	CommandOutputs []InstructionOutput `json:"commandOutputs" xml:"commandOutputs" form:"commandOutputs"`
	Error          *string             `json:"error" xml:"error" form:"error"`
}

var (
	defaultPanicMessage       = "wunderDB panicked in MAINTENANCE_MODE"
	defaultMaintenanceMessage = "wunderDB is running in MAINTENANCE_MODE"
	defaultMaintenancePort    = ":8081"
)

type MaintenanceModeOpts struct {
	Users  map[model.Identifier]*model.User
	Config config.Config
	app    *fiber.App
}

func NewMaintenanceMode(c config.Config) (*MaintenanceModeOpts, error) {
	users, error := LoadUsers(c)
	if error != nil {
		return nil, fmt.Errorf("error loading wfs: %s", error)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // fiber startup-message disable
	})

	return &MaintenanceModeOpts{
		Users:  users,
		Config: c,
		app:    app,
	}, nil
}

func (mmo *MaintenanceModeOpts) EnterMaintenanceMode() error {

	runModeStartupMessage(defaultMaintenanceMessage, defaultMaintenancePort)

	recoveryConf := recovery.DefaultConfig
	recoveryConf.Message = &defaultPanicMessage

	mmo.app.Use(logger.New())
	mmo.app.Use(recovery.New(recoveryConf))

	api := mmo.app.Group("/maintenance")

	// Maintenance Routes
	api.Post("/patch", mmo.Patch)
	api.Post("/close", mmo.Close)

	err := mmo.app.Listen(defaultMaintenancePort)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}

	return nil
}

func getErrorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func (mmo *MaintenanceModeOpts) Close(c *fiber.Ctx) error {
	if apiResponse, hasError := mmo.handleAuth(c); hasError {
		return mmo.sendResponse(c, *apiResponse, !shutdown)
	}

	return mmo.sendResponse(c, PatchApiResponse{
		Status:         "success",
		Message:        "maintenance mode shutting down",
		CommandOutputs: []InstructionOutput{},
		Error:          nil,
	}, shutdown)
}

func (mmo *MaintenanceModeOpts) handleAuth(c *fiber.Ctx) (*PatchApiResponse, bool) {
	patchApiResponse := PatchApiResponse{}

	username, password, err := authentication.HandleUserCredentials(c.Get(Authorization))
	if err != nil {
		patchApiResponse.Status = "unauthorized"
		patchApiResponse.Message = "error parsing credentials"
		patchApiResponse.Error = &err.ErrMessage
		patchApiResponse.CommandOutputs = []InstructionOutput{}
		return &patchApiResponse, true
	}

	isAuthorized, err := mmo.AuthenticateUser(model.Identifier(*username), *password)
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

func (mmo *MaintenanceModeOpts) Patch(c *fiber.Ctx) error {
	patchApiResponse := PatchApiResponse{}

	if apiResponse, hasError := mmo.handleAuth(c); hasError {
		return mmo.sendResponse(c, *apiResponse, !shutdown)
	}

	instructions := new(Instruction)
	if err := c.BodyParser(instructions); err != nil {
		return err
	}

	cmdOutputs := []InstructionOutput{}
	for _, command := range instructions.Commands {
		stdout, err := exec.Command("sh", "-c", command).Output()

		cmdOutput := InstructionOutput{
			Command:   command,
			Output:    string(stdout),
			Error:     getErrorString(err),
			IsSuccess: err == nil,
		}

		cmdOutputs = append(cmdOutputs, cmdOutput)
	}

	patchApiResponse.Status = "success"
	patchApiResponse.CommandOutputs = cmdOutputs
	patchApiResponse.Message = "commands executed successfully"
	patchApiResponse.Error = nil

	return mmo.sendResponse(c, patchApiResponse, false)
}

func LoadUsers(c config.Config) (map[model.Identifier]*model.User, error) {
	fs := wfs.NewWFileSystem(c.PersistantStoragePath)

	loadedUsers, err := fs.LoadUsers()
	if err != nil {
		return nil, fmt.Errorf("error loading wfs: %s", err)
	}

	return loadedUsers, nil
}

func (mmo *MaintenanceModeOpts) AuthenticateUser(requesterId model.Identifier, password string) (bool, *wdbErrors.WdbError) {
	if requesterId != model.Identifier(mmo.Config.AdminID) {
		return false, &wdbErrors.WdbError{
			ErrCode:        "unauthorizedInMaintenanceMode",
			ErrMessage:     "only admin can authenticate in maintenance mode",
			HttpStatusCode: fiber.StatusUnauthorized,
		}
	}

	user, exists := mmo.Users[requesterId]
	if !exists {
		return authentication.InvalidUser, &wdbErrors.AuthenticatingUserDoesNotExist
	}

	hashedPassword := authentication.Hash(password, user.Authentication.HashingAlgorithm)
	if user.Authentication.HashedSecret == hashedPassword {
		return authentication.ValidUser, nil
	}

	return authentication.InvalidUser, &wdbErrors.InvalidCredentialsError
}

func (mmo *MaintenanceModeOpts) sendResponse(c *fiber.Ctx, apiResponse PatchApiResponse, shutdown bool) error {
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
			mmo.app.Shutdown()
		}()
	}

	return nil
}
