package backend

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/controller"
	"go-skeleton/internal/request/user"
	"go-skeleton/pkg/response"
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
	var params = &user.IndexReq{}
	if err := c.Validate(ctx, params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJson(ctx, "", "index")
}
