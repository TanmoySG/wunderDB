package recovery

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

var (
	defaultMessage string = "panic on request"
)

type recoveryMessage struct {
	Message *string `json:"message,omitempty"`
	Stack   *string `json:"stack,omitempty"`
}

func (rm recoveryMessage) Marshal() []byte {
	responseJsonBytes, err := json.Marshal(rm)
	if err != nil {
		return nil
	}
	return responseJsonBytes
}

// Config defines the config for recovery middleware.
type Config struct {
	Next             func(c *fiber.Ctx) bool
	SendMessage      bool
	Message          *string
	EnableStackTrace bool
	RecoveryHandler  func(c *fiber.Ctx)
}

// ConfigDefault is the default config
var DefaultConfig = Config{
	Next:             nil,
	SendMessage:      true,
	Message:          &defaultMessage,
	EnableStackTrace: true,
	RecoveryHandler:  nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return DefaultConfig
	}

	// Override default config
	cfg := config[0]

	return cfg
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				if cfg.SendMessage {
					var stackTrace *string

					if cfg.EnableStackTrace {
						stackTraceString := string(debug.Stack())
						stackTrace = &stackTraceString
					}

					r := recoveryMessage{
						Message: cfg.Message,
						Stack:   stackTrace,
					}
					c.Status(http.StatusInternalServerError)
					c.Send(r.Marshal())

					cfg.RecoveryHandler(c)
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
