package handlers

import (
	"fmt"

	er "github.com/TanmoySG/wunderDB/internal/errors"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	w "github.com/TanmoySG/wunderDB/pkg/wdb"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	CheckPermissions(c *fiber.Ctx) error
	GrantRoles(c *fiber.Ctx) error
}

func NewHandlers(client w.Client) Client {
	return wdbHandlers{
		wdbClient: client,
	}
}

func (wh wdbHandlers) Hello(c *fiber.Ctx) error {
	msg := map[string]string{
		"message": fmt.Sprintf("✋ %s", "hello"),
	}
	resp := response.Format("ping", nil, msg)
	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}

func SendResponse(c *fiber.Ctx, marshaledResponse []byte, statusCode int) error {
	c.Set(ContentType, ApplicationJson)
	c.Send(marshaledResponse)
	return c.SendStatus(statusCode)
}

func ValidateRequest(request any) *er.WdbError {
	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		return &er.ValidationError
	}
	return nil
}
