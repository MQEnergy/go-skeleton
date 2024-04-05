package routes

import (
	"github.com/MQEnergy/go-skeleton/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func InitFrontendGroup(r fiber.Router, handles ...fiber.Handler) {
	router := r.Group("api", handles...)
	{
		router.Get("/", func(ctx *fiber.Ctx) error {
			return response.SuccessJSON(ctx, "", "api")
		})
	}
}
