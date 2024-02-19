package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/pkg/validator"
	"go-skeleton/pkg/response"
)

type Controller struct{}

var Base = &Controller{}

func (c *Controller) Index(ctx *fiber.Ctx) error {
	return response.SuccessJson(ctx, "index", "")
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	return response.SuccessJson(ctx, "create", "")
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	return response.SuccessJson(ctx, "delete", "")
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	return response.SuccessJson(ctx, "update", "")
}

func (c *Controller) View(ctx *fiber.Ctx) error {
	return response.SuccessJson(ctx, "view", "")
}

// ValidateReqParams 验证请求参数
func (c *Controller) ValidateReqParams(ctx *fiber.Ctx, requestParams fiber.Map) error {
	var err error
	switch string(ctx.Request().Header.ContentType()) {
	case "application/json":
		err = ctx.Bind(requestParams)
	case "application/xml":
		err = ctx.XML(requestParams)
	default:
		err = ctx.Bind(requestParams)
	}
	if err != nil {
		translate := validator.Translate(err)
		return errors.New(translate[0])
	}
	return nil
}
