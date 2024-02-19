package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-skeleton/internal/middleware"
	"go-skeleton/internal/router/routes"
)

// 路由分组
var (
	publicMiddleware = []fiber.Handler{
		middleware.IpAuth,
	}
)

// Register ...
func Register() *fiber.App {
	r := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})
	// middleware cors, cache, X-Request-Id
	r.Use(cors.New(), cache.New(), requestid.New())

	// common
	routes.InitCommonGroup(r, publicMiddleware...)
	// backend
	routes.InitBackendGroup(r, middleware.CasbinAuth)
	// frontend
	routes.InitFrontendGroup(r, middleware.CasbinAuth)

	return r
}
