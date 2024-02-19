package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JsonResponse struct {
	Status    int         `json:"status"`
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

	return ctx.JSON(JsonResponse{
		Status:    status,
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
	return Json(ctx, http.StatusOK, Success, message, data)
}

// BadRequestException 400错误
func BadRequestException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	return Json(ctx, http.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	return Json(ctx, http.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[Failed]
	}
	return Json(ctx, http.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	return Json(ctx, http.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *fiber.Ctx, message string) error {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	return Json(ctx, http.StatusInternalServerError, InternalErr, message, nil)
}
