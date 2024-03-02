package router

import (
	"go-skeleton/internal/middleware"
	"go-skeleton/internal/router/routes"
	"go-skeleton/pkg/response"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Register ...
func Register(appName string) *fiber.App {
	publicMiddleware := []fiber.Handler{
		middleware.LoggerMiddleware(),  // 日志
		middleware.WhiteIpMiddleware(), // 白名单
	}

	r := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.NotFoundException(c, err.Error())
		},
		DisableStartupMessage: false,        // When set to true, it will not print out debug information
		AppName:               appName,      // This allows to setup app name for the app
		JSONEncoder:           json.Marshal, // If you're not happy with the performance of encoding/json, we recommend you to use these libraries
		JSONDecoder:           json.Unmarshal,
	})
	// middleware cors, compress, cache, X-Request-Id
	r.Use(
		recover.New(),
		cors.New(),
		compress.New(),
		requestid.New(),
	)
	// common
	routes.InitCommonGroup(r, publicMiddleware...)
	// backend
	routes.InitBackendGroup(r, middleware.CasbinMiddleware(), middleware.AuthMiddleware(), middleware.CacheMiddleware())
	// frontend
	routes.InitFrontendGroup(r, append(publicMiddleware, middleware.AuthMiddleware())...)

	return r
}
