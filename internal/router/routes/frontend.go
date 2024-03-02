package routes

import (
	"github.com/MQEnergy/go-skeleton/internal/app/controller/backend"

	"github.com/gofiber/fiber/v2"
)

func InitFrontendGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("api", middleware...)
	{
		router.Post("/user/index", backend.User.Index)
	}
}
