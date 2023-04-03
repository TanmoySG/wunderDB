package server

import (
	"fmt"
	"log"

	"github.com/TanmoySG/wunderDB/internal/server/handlers"
	"github.com/TanmoySG/wunderDB/internal/server/middlewares/recovery"
	"github.com/TanmoySG/wunderDB/internal/server/routes"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	defaultPanicMessage = "wunderDB panicked on request"
)

type wdbServer struct {
	port    string
	handler handlers.Client
}

type Client interface {
	Start()
}

func NewWdbServer(wdbClient wdbClient.Client, wdbBasePath, port string) Client {
	return wdbServer{
		port:    fmt.Sprintf(":%s", port),
		handler: handlers.NewHandlers(wdbClient, wdbBasePath),
	}
}

func (ws wdbServer) Start() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // fiber box disable
	})

	// recovery configuration
	recoveryConf := recovery.DefaultConfig
	recoveryConf.Message = &defaultPanicMessage

	// transaction logging configuration
	// transactionLoggingConfig := transactionLogging.DefaultConfig
	// transactionLoggingConfig.Skip = func(c *fiber.Ctx) bool {
	// 	return false
	// }

	app.Use(logger.New())
	// app.Use(recovery.New(recoveryConf))
	// app.Use(transactionLogging.New(transactionLoggingConfig))

	api := app.Group("/api")

	// api home route
	api.Get("/", ws.handler.Hello)

	// Database Routes
	api.Post(routes.CreateDatabase, ws.handler.CreateDatabase)
	api.Get(routes.FetchDatabase, ws.handler.FetchDatabase)
	api.Delete(routes.DeleteDatabase, ws.handler.DeleteDatabase)

	// Collection Routes
	api.Post(routes.CreateCollection, ws.handler.CreateCollection)
	api.Get(routes.FetchCollection, ws.handler.FetchCollection)
	api.Delete(routes.DeleteCollection, ws.handler.DeleteCollection)

	// Data Routes
	api.Post(routes.AddData, ws.handler.AddData)
	api.Get(routes.ReadData, ws.handler.ReadData)
	api.Delete(routes.DeleteData, ws.handler.DeleteData)
	api.Patch(routes.UpdateData, ws.handler.UpdateData)

	// Role Routes
	api.Post(routes.CreateRole, ws.handler.CreateRole)
	api.Get(routes.ListRoles, ws.handler.ListRoles)

	// User Routes
	api.Post(routes.CreateUser, ws.handler.CreateUser)
	api.Post(routes.GrantRoles, ws.handler.GrantRoles)
	api.Get(routes.LoginUser, ws.handler.LoginUser)
	// app.Get("/api/users/permission", ws.handler.CheckPermissions)

	err := app.Listen(ws.port)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}
}

// func getRequestAction(c fiber.Ctx) string {
// 	var responseBody map[string]interface{}

// 	responseBodyBytes := c.Response().Body()

// 	err := json.Unmarshal(responseBodyBytes, &responseBody)
// 	if err != nil {
// 		return ""
// 	}

// 	if _, ok := responseBody["action"]; ok {
// 		return responseBody["action"].(string)
// 	}

// 	return ""
// }
