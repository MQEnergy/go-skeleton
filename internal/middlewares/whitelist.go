package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/helper"
	"go-skeleton/pkg/response"
	"strings"
)

// WhiteIpMiddleware white list middleware
func WhiteIpMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientIp := ctx.IP()
		flag := false
		whiteList := variable.Config.GetString("server.whiteList")
		if whiteList == "*" || whiteList == "" || ctx.IsFromLocal() {
			flag = true
		} else {
			if helper.InAnySlice(strings.Split(whiteList, ","), clientIp) {
				flag = true
			}
		}
		if !flag {
			return response.UnauthorizedException(ctx, fmt.Sprintf("%s is not in the white list", clientIp))
		}
		return ctx.Next()
	}
}
