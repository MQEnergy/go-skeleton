package routes

import (
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r fiber.Router, handles ...fiber.Handler) {
	router := r.Group("backend", handles...)
	{
		router.Get("/", func(ctx *fiber.Ctx) error {
			return response.SuccessJSON(ctx, "", "backend")
		})
	}
}
