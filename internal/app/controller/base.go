package controller

import (
	"errors"
	"github.com/MQEnergy/go-skeleton/internal/app/pkg/validator"
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"reflect"

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
	dest := reflect.ValueOf(param).Elem()
	paramType := dest.Type()

	queryInstance := reflect.New(paramType).Interface()
	if err := ctx.QueryParser(queryInstance); err != nil {
		return errors.New("参数格式错误")
	}

	if ctx.Method() == fiber.MethodPost {
		if err := ctx.BodyParser(param); err != nil {
			return errors.New("参数格式错误")
		}
	}

	queryValue := reflect.ValueOf(queryInstance).Elem()
	for i := 0; i < dest.NumField(); i++ {
		if _, hasQuery := paramType.Field(i).Tag.Lookup("query"); hasQuery {
			dest.Field(i).Set(queryValue.Field(i))
		}
	}

	translates := validate.Validate(param)
	if len(translates) > 0 && translates[0].Error {
		return errors.New(translates[0].Message)
	}
	return nil
}
