package response

import (
	"github.com/gofiber/fiber/v2"
)

type JsonResponse struct {
	ErrCode   Code        `json:"errcode"`
	RequestId string      `json:"requestid"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// Json 基础返回
func Json(ctx *fiber.Ctx, status int, errcode Code, message string, data interface{}) error {
	if message == "" {
		message = CodeMap[errcode]
	}
	return ctx.Status(status).JSON(JsonResponse{
		ErrCode:   errcode,
		Message:   message,
		RequestId: ctx.GetRespHeader("X-Request-Id"),
		Data:      data,
	})
}

// SuccessJson 成功返回
func SuccessJson(ctx *fiber.Ctx, message string, data interface{}) error {
	if message == "" {
		message = Success.Msg()
	}
	return Json(ctx, fiber.StatusOK, Success, message, data)
}

// BadRequestException 400错误
func BadRequestException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	return Json(ctx, fiber.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	return Json(ctx, fiber.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[Failed]
	}
	return Json(ctx, fiber.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	return Json(ctx, fiber.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	return Json(ctx, fiber.StatusInternalServerError, InternalErr, message, nil)
}
