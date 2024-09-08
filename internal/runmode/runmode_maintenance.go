package runmode

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/TanmoySG/wunderDB/internal/config"
	"github.com/TanmoySG/wunderDB/internal/server/middlewares/recovery"
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
	defaultMaintenanceMessage = "wunderDB is running in "
	defaultMaintenancePort    = ":8086"
)

type MaintenanceModeOpts struct {
	RunModeOpts RunModeOpts
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
		RunModeOpts{
			Users:  users,
			Config: c,
			app:    app,
		},
	}, nil
}

func (mmo *MaintenanceModeOpts) EnterMaintenanceMode() error {

	runModeStartupMessage(
		string(RUN_MODE_MAINTENANCE),
		defaultMaintenancePort,
	)

	recoveryConf := recovery.DefaultConfig
	recoveryConf.Message = &defaultPanicMessage

	mmo.RunModeOpts.app.Use(logger.New())
	mmo.RunModeOpts.app.Use(recovery.New(recoveryConf))

	api := mmo.RunModeOpts.app.Group("/maintenance")

	// Maintenance Routes
	api.Post("/patch", mmo.Patch)
	api.Post("/close", mmo.Close)

	err := mmo.RunModeOpts.app.Listen(defaultMaintenancePort)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}

	return nil
}

func (mmo *MaintenanceModeOpts) Close(c *fiber.Ctx) error {
	if apiResponse, hasError := mmo.RunModeOpts.handleAuth(c); hasError {
		return mmo.RunModeOpts.sendResponse(c, *apiResponse, !shutdown)
	}

	return mmo.RunModeOpts.sendResponse(c, PatchApiResponse{
		Status:         "success",
		Message:        "maintenance mode shutting down",
		CommandOutputs: []InstructionOutput{},
		Error:          nil,
	}, shutdown)
}

func (mmo *MaintenanceModeOpts) Patch(c *fiber.Ctx) error {
	patchApiResponse := PatchApiResponse{}

	if apiResponse, hasError := mmo.RunModeOpts.handleAuth(c); hasError {
		return mmo.RunModeOpts.sendResponse(c, *apiResponse, !shutdown)
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

	return mmo.RunModeOpts.sendResponse(c, patchApiResponse, false)
}
