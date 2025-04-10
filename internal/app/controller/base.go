package controller

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MQEnergy/go-skeleton/internal/app/pkg/validator"
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

var Base = &Controller{}

func (c *Controller) Index(ctx *fiber.Ctx) error {
	return response.SuccessJSON(ctx, "index", "")
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	return response.SuccessJSON(ctx, "create", "")
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	return response.SuccessJSON(ctx, "delete", "")
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	return response.SuccessJSON(ctx, "update", "")
}

func (c *Controller) View(ctx *fiber.Ctx) error {
	return response.SuccessJSON(ctx, "view", "")
}

var validate validator.XValidator

func init() {
	valid, err := validator.New("zh")
	if err != nil {
		panic(err)
	}
	validate = *valid
}

// Validate ...
func (c *Controller) Validate(ctx *fiber.Ctx, param any) error {
	if err := ctx.ParamsParser(param); err != nil {
		return fmt.Errorf("路由参数解析失败: %v", err)
	}

	if ctx.Method() == fiber.MethodPost {
		contentType := string(ctx.Request().Header.ContentType())
		contentType = utils.ParseVendorSpecificContentType(contentType)
		ctypeEnd := strings.IndexByte(contentType, ';')
		if ctypeEnd != -1 {
			contentType = contentType[:ctypeEnd]
		}
		switch {
		case
			strings.HasPrefix(contentType, fiber.MIMEApplicationForm),
			strings.HasSuffix(contentType, "json"),
			strings.HasPrefix(contentType, fiber.MIMEMultipartForm),
			strings.HasPrefix(contentType, fiber.MIMETextXML),
			strings.HasPrefix(contentType, fiber.MIMEApplicationXML):
			if err := ctx.QueryParser(param); err != nil {
				return fmt.Errorf("查询参数解析失败: %v", err)
			}
			if err := ctx.BodyParser(param); err != nil {
				return fmt.Errorf("查询参数解析失败: %v", err)
			}

		default:
			if err := ctx.QueryParser(param); err != nil {
				return fmt.Errorf("查询参数解析失败: %v", err)
			}
		}
	} else {
		if err := ctx.QueryParser(param); err != nil {
			return fmt.Errorf("查询参数解析失败: %v", err)
		}
	}

	if translates := validate.Validate(param); len(translates) > 0 && translates[0].Error {
		return errors.New(translates[0].Message)
	}

	return nil
}
