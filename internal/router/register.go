package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-skeleton/internal/middleware"
	"go-skeleton/internal/router/routes"
	"go-skeleton/internal/variable"
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
		DisableStartupMessage: true,                                        // 设置为true时，不会打印出调试信息
		AppName:               variable.Config.GetString("server.appName"), // This allows to setup app name for the app
		//EnablePrintRoutes: true, // EnablePrintRoutes enables print all routes with their method, path, name and handler..
	})
	// middleware cors, cache, X-Request-Id
	r.Use(cors.New(), cache.New(), requestid.New())

	// logger
	if variable.Config.GetString("server.mode") != "production" {
		r.Use(logger.New(logger.Config{TimeFormat: "2006-01-02 15:04:05"}))
	}
	// common
	routes.InitCommonGroup(r, publicMiddleware...)
	// backend
	routes.InitBackendGroup(r, middleware.CasbinAuth)
	// frontend
	routes.InitFrontendGroup(r, middleware.CasbinAuth)

	return r
}
