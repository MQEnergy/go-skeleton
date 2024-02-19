package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/variable"
	"go-skeleton/pkg/utils"
	"net/http"
	"strings"
)

// IpAuth 白名单验证
func IpAuth(ctx *fiber.Ctx) error {
	clientIp := ctx.IP()
	flag := false
	whiteList := variable.Config.GetString("server.whiteList")
	if whiteList == "*" || whiteList == "" {
		flag = true
	} else {
		if utils.InAnySlice(strings.Split(whiteList, ","), clientIp) {
			flag = true
		}
	}
	if !flag {
		ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%s 不在ip白名单中", clientIp))
		return errors.New(fmt.Sprintf("%s 不在ip白名单中", clientIp))
	}
	return ctx.Next()
}
