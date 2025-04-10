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

// Index 用户列表
//
//	@Summary		用户列表
//	@Description	获取用户列表接口
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		user.IndexReq	true	"请求参数"
//	@Success		200		{object}  	response.JSONResponse  "成功"
//	@Failure		400		{object}  	response.JSONResponse  "请求错误"
//	@Router			/backend/user/index [get]
func (c *UserController) Index(ctx *fiber.Ctx) error {
	params := &user.IndexReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "index")
}

// Login 用户登录
//
//	@Summary		用户登录
//	@Description	用户登录接口
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		user.LoginReq	true	"登录请求参数"
//	@Success		200		{object}  	response.JSONResponse  "成功"
//	@Failure		400		{object}  	response.JSONResponse  "请求错误"
//	@Router			/backend/auth/login [post]
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
