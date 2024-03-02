package backend

import (
	"github.com/MQEnergy/go-skeleton/internal/app/controller"
	"github.com/MQEnergy/go-skeleton/internal/app/service/backend"
	"github.com/MQEnergy/go-skeleton/internal/request/user"
	"github.com/MQEnergy/go-skeleton/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	controller.Controller
}

var User = &UserController{}

// Index
// @Description: index
// @receiver c
// @param ctx
// @return error
// @author cx
func (c *UserController) Index(ctx *fiber.Ctx) error {
	params := &user.IndexReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "index")
}

// Login
// @Description: login
// @receiver c
// @param ctx
// @return error
// @author cx
func (c *UserController) Login(ctx *fiber.Ctx) error {
	params := &user.LoginReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := backend.Auth.Login(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}
