package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/controller/backend"
)

func InitFrontendGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("api", middleware...)
	{
		router.Post("/user/index", backend.User.Index)
		router.Post("/user/create", backend.User.Create)
		router.Get("/user/view", backend.User.View)
		router.Post("/user/update", backend.User.Update)
		router.Post("/user/delete", backend.User.Delete)
	}
}
