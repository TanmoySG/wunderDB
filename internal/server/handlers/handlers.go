package handlers

import (
	tx "github.com/TanmoySG/wunderDB/internal/txlogs"
	w "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/gofiber/fiber/v2"
)

type wdbHandlers struct {
	wdbClient w.Client
	wdbTxLogs tx.DotTxLog
	notices   *[]string
}

func NewHandlers(client w.Client, wdbBasePath string, notices ...string) Client {
	return wdbHandlers{
		wdbClient: client,
		wdbTxLogs: tx.UseDotTxLog(wdbBasePath),
		notices:   &notices,
	}
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
	QueryData(c *fiber.Ctx) error
	DeleteData(c *fiber.Ctx) error
	UpdateData(c *fiber.Ctx) error

	// Role Handlers
	ListRoles(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error

	// User Handlers
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	GrantRole(c *fiber.Ctx) error
	RevokeRole(c *fiber.Ctx) error
}
