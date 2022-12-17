package handlers

import (
	"fmt"

	w "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
)

const (
	CreateDatabaseAction = "create-database"
	FetchDatabaseAction  = "fetch-database"
	DeleteDatabaseAction = "delete-database"

	CreateCollectionAction = "create-collection"
	FetchCollectionAction  = "fetch-collection"
	DeleteCollectionAction = "delete-collection"
)

type wdbHandlers struct {
	wdbClient w.Client
}

type Client interface {
	Hello(c *fiber.Ctx) error

	// Database Handlers
	CreateDatabase(c *fiber.Ctx) error
	FetchDatabase(c *fiber.Ctx) error
	DeleteDatabase(c *fiber.Ctx) error

	// Collection Handlers
	CreateCollection(c *fiber.Ctx) error
	FetchCollection(c *fiber.Ctx) error
	DeleteCollection(c *fiber.Ctx) error
}

func NewHandlers(client w.Client) Client {
	return wdbHandlers{
		wdbClient: client,
	}
}

func (wh wdbHandlers) Hello(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", "hello")
	return c.SendString(msg) // => ✋ register
}
