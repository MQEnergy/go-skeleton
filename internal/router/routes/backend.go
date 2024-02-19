package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/controller/backend"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("backend", middleware...)
	{
		router.Post("/user/create", backend.User.Create)
		router.Get("/user/view", backend.User.View)
		router.Post("/user/update", backend.User.Update)
		router.Post("/user/delete", backend.User.Delete)
	}
}
