package handlers

import (
	"fmt"
	"os"

	"github.com/TanmoySG/wunderDB/internal/privileges"
	"github.com/TanmoySG/wunderDB/internal/server/response"
	"github.com/TanmoySG/wunderDB/internal/version"
	"github.com/gofiber/fiber/v2"
)

func (wh wdbHandlers) Hello(c *fiber.Ctx) error {
	privilege := privileges.Ping
	ua := c.Get("User-Agent")

	msg := map[string]interface{}{
		"wunderDb": map[string]interface{}{
			"version":    version.WDB_VERSION,
			"build-date": version.BUILD_DATE,
			"notice":     os.Getenv("NOTICE"),
		},
		"message":    fmt.Sprintf("✋ %s", "hello"),
		"user-agent": ua,
	}

	resp := response.Format(privilege, nil, msg, *wh.notices...)
	if err := sendResponse(c, resp); err != nil {
		return err
	}

	wh.handleTransactions(c, resp, noEntities)
	return nil
}
