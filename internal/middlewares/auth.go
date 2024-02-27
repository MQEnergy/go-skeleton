package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/response"
)

// AuthMiddleware jwt authentication middleware
func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(vars.Config.GetString("jwt.secret"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.UnauthorizedException(c, err.Error())
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user, ok := ctx.Locals("user").(*jwt.Token)
			if ok {
				if claims, ok := user.Claims.(jwt.MapClaims); ok && user.Valid {
					if sub, ok := claims["sub"].(map[string]interface{}); ok {
						ctx.Set("uuid", sub["uuid"].(string))
						ctx.Set("uid", gconv.String(sub["id"]))
						ctx.Set("role_ids", sub["role_ids"].(string))
						return ctx.Next()
					}
				}
			}
			return response.UnauthorizedException(ctx, "token is invalid")
		},
		Filter: func(ctx *fiber.Ctx) bool {
			if ctx.Path() == "/backend/auth/login" {
				return true
			}
			return false
		},
		// ContextKey: "user", // used in ctx.Locals("user").(*jwt.Token)
	})
}
