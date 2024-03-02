package middleware

import (
	"fmt"
	"strings"

	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/helper"
	"github.com/MQEnergy/go-skeleton/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// WhiteIpMiddleware white list middleware
func WhiteIpMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientIp := ctx.IP()
		flag := false
		whiteList := vars.Config.GetString("server.whiteList")
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
