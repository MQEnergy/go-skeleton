package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/pkg/response"
)

func InitCommonGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("/", middleware...)
	{
		router.Get("/", func(c *fiber.Ctx) error {
			return response.SuccessJson(c, "", "")
		})
		router.Get("/ping", func(c *fiber.Ctx) error {
			return response.SuccessJson(c, "", "pong")
		})
	}
}
