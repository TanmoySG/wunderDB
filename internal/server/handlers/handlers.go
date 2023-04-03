package handlers

import (
	tx "github.com/TanmoySG/wunderDB/internal/txlogs"
	w "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
)

type wdbHandlers struct {
	wdbClient w.Client
	wdbTxLogs tx.DotTxLog
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

	// Data Handlers
	AddData(c *fiber.Ctx) error
	ReadData(c *fiber.Ctx) error
	DeleteData(c *fiber.Ctx) error
	UpdateData(c *fiber.Ctx) error

	// Role Handlers
	CreateRole(c *fiber.Ctx) error
	ListRoles(c *fiber.Ctx) error

	// User Handlers
	CreateUser(c *fiber.Ctx) error
	GrantRoles(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	// CheckPermissions(c *fiber.Ctx) error

}

func NewHandlers(client w.Client) Client {
	return wdbHandlers{
		wdbClient: client,
	}
}
