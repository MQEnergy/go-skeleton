package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/response"
	"go-skeleton/pkg/utils"
	"strings"
)

// IpAuth 白名单验证
func IpAuth(ctx *fiber.Ctx) error {
	clientIp := ctx.IP()
	flag := false
	whiteList := variable.Config.GetString("server.whiteList")
	if whiteList == "*" || whiteList == "" || ctx.IsFromLocal() {
		flag = true
	} else {
		if utils.InAnySlice(strings.Split(whiteList, ","), clientIp) {
			flag = true
		}
	}
	if !flag {
		return response.UnauthorizedException(ctx, fmt.Sprintf("%s 不在ip白名单中", clientIp))
	}
	return ctx.Next()
}
