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
	response.Json(ctx, fiber.StatusOK, response.Success, "index", "")
	return ctx.Next()
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	response.Json(ctx, fiber.StatusOK, response.Success, "create", "")
	return ctx.Next()
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	response.Json(ctx, fiber.StatusOK, response.Success, "delete", "")
	return ctx.Next()
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	response.Json(ctx, fiber.StatusOK, response.Success, "update", "")
	return ctx.Next()
}

func (c *Controller) View(ctx *fiber.Ctx) error {
	response.Json(ctx, fiber.StatusOK, response.Success, "view", "")
	return ctx.Next()
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
