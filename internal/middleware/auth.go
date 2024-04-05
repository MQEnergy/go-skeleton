package middleware

import (
	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/response"
	"github.com/spf13/cast"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware jwt authentication middleware
func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(vars.Config.GetString("jwt.secret"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.UnauthorizedException(c, "会话已过期，请重新登录 err: "+err.Error())
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			if user, ok := ctx.Locals("user").(*jwt.Token); ok {
				if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
					if sub, ok := claims["sub"].(map[string]interface{}); ok {
						ctx.Set("uuid", cast.ToString(sub["uuid"]))
						ctx.Set("uid", cast.ToString(sub["id"]))
						ctx.Set("role_ids", cast.ToString(sub["role_ids"]))
						return ctx.Next()
					}
				}
			}
			return response.UnauthorizedException(ctx, "token is invalid")
		},
		Filter: func(ctx *fiber.Ctx) bool {
			// notice: 在此可自定义路由通过auth验证
			//if strings.HasPrefix(ctx.Path(), "/backend/auth/login") {
			//	return true
			//}
			return false
		},
		// ContextKey: "user", // used in ctx.Locals("user").(*jwt.Token)
	})
}
