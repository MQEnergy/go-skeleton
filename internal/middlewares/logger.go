package middlewares

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/logger"
)

// LoggerMiddleware ...
func LoggerMiddleware() fiber.Handler {
	return slogfiber.NewWithConfig(logger.New(
		vars.Config,
	), slogfiber.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
		WithUserAgent:    true,
		WithRequestBody:  true,
		WithRequestID:    true,
	})
}
