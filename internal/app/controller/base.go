package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/app/pkg/validator"
	"go-skeleton/pkg/response"
	"strings"
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

var validate *validator.XValidator

func init() {
	var err error
	validate, err = validator.New("zh")
	if err != nil {
		panic(err)
	}
}

// Validate ...
func (c *Controller) Validate(ctx *fiber.Ctx, param any) error {
	var errs = make([]error, 0)
	// post
	if ctx.Method() == fiber.MethodPost {
		contentType := string(ctx.Request().Header.ContentType())
		switch {
		case
			contentType == "application/x-www-form-urlencoded",
			contentType == "multipart/form-data",
			contentType == "application/json",
			contentType == "application/xml",
			contentType == "text/xml",
			strings.Contains(contentType, "multipart/form-data") == true:
			if err := ctx.BodyParser(param); err != nil {
				errs = append(errs, err)
			}
			if err := ctx.QueryParser(param); err != nil {
				errs = append(errs, err)
			}
		}
	}
	// get
	if ctx.Method() == fiber.MethodGet {
		if err := ctx.QueryParser(param); err != nil {
			errs = append(errs, err)
		}
	}
	// path
	if err := ctx.ParamsParser(param); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs[0]
	}
	translates := validate.Validate(param)
	if len(translates) > 0 && translates[0].Error {
		return errors.New(translates[0].Message)
	}
	return nil
}
