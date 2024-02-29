package routes

import (
	"go-skeleton/internal/app/controller/backend"

	"github.com/gofiber/fiber/v2"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("backend", middleware...)
	{
		router.Post("/auth/login", backend.User.Login)
	}
}
