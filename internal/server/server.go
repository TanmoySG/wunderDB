package server

import (
	"fmt"
	"log"

	"github.com/TanmoySG/wunderDB/internal/server/handlers"
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
	app.Post("/api/databases", ws.handler.CreateDatabase)
	app.Get("/api/databases/:database", ws.handler.FetchDatabase)
	app.Delete("/api/databases/:database", ws.handler.DeleteDatabase)

	// Collection Routes
	app.Post("/api/databases/:database/collections", ws.handler.CreateCollection)
	app.Get("/api/databases/:database/collections/:collection", ws.handler.FetchCollection)
	app.Delete("/api/databases/:database/collections/:collection", ws.handler.DeleteCollection)

	// Data Routes
	app.Post("/api/databases/:database/collections/:collection/data", ws.handler.AddData)
	app.Get("/api/databases/:database/collections/:collection/data", ws.handler.ReadData)
	app.Delete("/api/databases/:database/collections/:collection/data", ws.handler.DeleteData)
	app.Patch("/api/databases/:database/collections/:collection/data", ws.handler.UpdateData)

	// Role Routes
	app.Post("/api/roles", ws.handler.CreateRole)
	app.Get("/api/roles", ws.handler.ListRoles)

	// User Routes
	app.Post("/api/users", ws.handler.CreateUser)
	app.Post("/api/users/grant", ws.handler.GrantRoles)
	// app.Get("/api/users/permission", ws.handler.CheckPermissions)

	err := app.Listen(ws.port)
	if err != nil {
		log.Fatalf("exiting wdb: %s", err)
	}
}
