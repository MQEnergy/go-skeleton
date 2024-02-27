package response

import (
	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	ErrCode   Code        `json:"errcode"`
	RequestID string      `json:"requestid"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// JSON 基础返回
func JSON(ctx *fiber.Ctx, status int, errcode Code, message string, data interface{}) error {
	if message == "" {
		message = CodeMap[errcode]
	}
	return ctx.Status(status).JSON(JSONResponse{
		ErrCode:   errcode,
		Message:   message,
		RequestID: ctx.GetRespHeader("X-Request-Id"),
		Data:      data,
	})
}

// SuccessJSON 成功返回
func SuccessJSON(ctx *fiber.Ctx, message string, data interface{}) error {
	if message == "" {
		message = Success.Msg()
	}
	return JSON(ctx, fiber.StatusOK, Success, message, data)
}

// BadRequestException 400错误
func BadRequestException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	return JSON(ctx, fiber.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	return JSON(ctx, fiber.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[Failed]
	}
	return JSON(ctx, fiber.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	return JSON(ctx, fiber.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	return JSON(ctx, fiber.StatusInternalServerError, InternalErr, message, nil)
}
