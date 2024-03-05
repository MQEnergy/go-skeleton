package routes

import (
	"github.com/MQEnergy/go-skeleton/internal/app/controller/backend"
	"github.com/gofiber/fiber/v2"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r fiber.Router, handles ...fiber.Handler) {
	router := r.Group("backend", handles...)
	{
		router.Post("/auth/login", backend.User.Login)
	}
}
