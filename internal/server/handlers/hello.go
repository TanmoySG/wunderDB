package handlers

import (
	"fmt"

	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/gofiber/fiber/v2"
)

func (wh wdbHandlers) Hello(c *fiber.Ctx) error {
	ua := c.Get("User-Agent")

	msg := map[string]string{
		"message":    fmt.Sprintf("✋ %s", "hello"),
		"user-agent": ua,
	}

	resp := response.Format("ping", nil, msg)
	return SendResponse(c, resp.Marshal(), resp.HttpStatusCode)
}
