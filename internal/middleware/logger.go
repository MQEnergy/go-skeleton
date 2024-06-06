package middleware

import (
	"log/slog"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/logger"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

// LoggerMiddleware ...
func LoggerMiddleware() fiber.Handler {
	return slogfiber.NewWithConfig(logger.New(
		vars.Config.GetString("log.fileName"),
		&vars.Config,
	), slogfiber.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
		WithUserAgent:    true,
		WithRequestBody:  true,
		WithRequestID:    true,
		Filters: []slogfiber.Filter{
			func(c *fiber.Ctx) bool {
				if file, err := c.FormFile("file"); err == nil {
					if file != nil {
						return false
					}
				}
				return true
			},
		},
	})
}
