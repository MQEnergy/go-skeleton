package router

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-skeleton/internal/middlewares"
	"go-skeleton/internal/router/routes"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/response"
	"time"
)

// Register ...
func Register(appName string) *fiber.App {
	var (
		publicMiddleware = []fiber.Handler{
			middlewares.WhiteIpMiddleware(),
		}
	)

	r := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.NotFoundException(c, err.Error())
		},
		DisableStartupMessage: true,         // When set to true, it will not print out debug information
		AppName:               appName,      // This allows to setup app name for the app
		JSONEncoder:           json.Marshal, // If you're not happy with the performance of encoding/json, we recommend you to use these libraries
		JSONDecoder:           json.Unmarshal,
	})
	// middleware cors, compress, cache, X-Request-Id
	r.Use(
		cors.New(),
		compress.New(),
		requestid.New(),
	)
	// logger
	if variable.Config.GetString("server.mode") != "production" {
		r.Use(logger.New(logger.Config{TimeFormat: time.DateTime}))
	}
	// common
	routes.InitCommonGroup(r, publicMiddleware...)
	// backend
	routes.InitBackendGroup(r, middlewares.CasbinMiddleware(), middlewares.AuthMiddleware(), middlewares.CacheMiddleware())
	// frontend
	routes.InitFrontendGroup(r, append(publicMiddleware, middlewares.AuthMiddleware())...)

	return r
}
