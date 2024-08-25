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

type wdbServer struct {
	port    string
	handler handlers.Client
	notices []string
}

type Client interface {
	Start()
}

func NewWdbServer(wdbClient wdbClient.Client, wdbBasePath, port string, notices ...string) Client {
	return wdbServer{
		port:    fmt.Sprintf(":%s", port),
		handler: handlers.NewHandlers(wdbClient, wdbBasePath, notices...),
		notices: notices,
	}
}

func (ws wdbServer) Start() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, // fiber startup-message disable
	})

	ws.startupMessage("", ws.port)

	// recovery configuration
	recoveryConf := recovery.DefaultConfig
	recoveryConf.Message = &defaultPanicMessage

	app.Use(logger.New())
	app.Use(recovery.New(recoveryConf))

	api := app.Group("/api")

	// API Home Route
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
	api.Post(routes.AddRecords, ws.handler.AddRecords)
	api.Get(routes.ReadRecords, ws.handler.ReadRecords)
	api.Post(routes.QueryRecords, ws.handler.QueryRecords)
	api.Delete(routes.DeleteRecords, ws.handler.DeleteRecords)
	api.Patch(routes.UpdateRecords, ws.handler.UpdateRecords)

	// Role Routes
	api.Post(routes.CreateRole, ws.handler.CreateRole)
	api.Get(routes.ListRoles, ws.handler.ListRoles)
	api.Patch(routes.UpdateRole, ws.handler.UpdateRole)

	// User Routes
	api.Post(routes.CreateUser, ws.handler.CreateUser)
	api.Post(routes.GrantRole, ws.handler.GrantRole)
	api.Get(routes.LoginUser, ws.handler.LoginUser)
	api.Delete(routes.RevokeRole, ws.handler.RevokeRole)

	err := app.Listen(ws.port)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}
}
