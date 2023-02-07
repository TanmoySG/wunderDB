package server

import (
	"fmt"
	"log"

	"github.com/TanmoySG/wunderDB/internal/server/handlers"
	"github.com/TanmoySG/wunderDB/internal/server/routes"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type wdbServer struct {
	port    string
	handler handlers.Client
}

type Client interface {
	Start()
}

func NewWdbServer(wdbClient wdbClient.Client, port string) Client {
	return wdbServer{
		port:    fmt.Sprintf(":%s", port),
		handler: handlers.NewHandlers(wdbClient),
	}
}

func (ws wdbServer) Start() {

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/api", ws.handler.Hello)

	// Database Routes
	app.Post(routes.CreateDatabase, ws.handler.CreateDatabase)
	app.Get(routes.FetchDatabase, ws.handler.FetchDatabase)
	app.Delete(routes.DeleteDatabase, ws.handler.DeleteDatabase)

	// Collection Routes
	app.Post(routes.CreateCollection, ws.handler.CreateCollection)
	app.Get(routes.FetchCollection, ws.handler.FetchCollection)
	app.Delete(routes.DeleteCollection, ws.handler.DeleteCollection)

	// Data Routes
	app.Post(routes.AddData, ws.handler.AddData)
	app.Get(routes.ReadData, ws.handler.ReadData)
	app.Delete(routes.DeleteData, ws.handler.DeleteData)
	app.Patch(routes.UpdateData, ws.handler.UpdateData)

	// Role Routes
	app.Post(routes.CreateRole, ws.handler.CreateRole)
	app.Get(routes.ListRoles, ws.handler.ListRoles)

	// User Routes
	app.Post(routes.CreateUser, ws.handler.CreateUser)
	app.Post(routes.GrantRoles, ws.handler.GrantRoles)
	app.Get(routes.LoginUser, ws.handler.LoginUser)
	// app.Get("/api/users/permission", ws.handler.CheckPermissions)

	err := app.Listen(ws.port)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}
}
