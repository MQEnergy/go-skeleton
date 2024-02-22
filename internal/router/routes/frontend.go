package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/controller/backend"
)

func InitFrontendGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("api", middleware...)
	{
		router.Post("/user/index", backend.User.Index)
	}
}
