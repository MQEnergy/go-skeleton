package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitCommonGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("/", middleware...)
	{
		router.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("/")
		})
		router.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})
	}
}
