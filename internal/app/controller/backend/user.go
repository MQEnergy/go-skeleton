package backend

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/controller"
)

type UserController struct {
	controller.Controller
}

var User = &UserController{}

func (c *UserController) Index(ctx *fiber.Ctx) error {
	return ctx.Next()
}
