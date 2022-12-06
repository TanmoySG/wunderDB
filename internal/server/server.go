package server

import (
	"github.com/TanmoySG/wunderDB/internal/server/handlers"
	wdbClient "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type wdbServer struct {
	handler handlers.Client
}

type Client interface {
	Start()
}

func NewWdbServer(wdbClient wdbClient.Client) Client {
	return wdbServer{
		handler: handlers.NewHandlers(wdbClient),
	}
}

func (ws wdbServer) Start() {

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/api", ws.handler.Hello)

	app.Post("/api/databases", ws.handler.CreateDatabase)
	app.Get("/api/databases/:database", ws.handler.FetchDatabase)
	app.Delete("/api/databases/:database", ws.handler.DeleteDatabase)

	app.Post("/api/databases/:database/collections", ws.handler.CreateCollection)
	app.Get("/api/databases/:database/collections/:collection", ws.handler.FetchCollection)
	app.Delete("/api/databases/:database/collections/:collection", ws.handler.DeleteCollection)

	app.Listen(":3000")
}
