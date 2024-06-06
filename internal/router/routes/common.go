package routes

import (
	"github.com/MQEnergy/go-skeleton/internal/app/controller/backend"
	"github.com/MQEnergy/go-skeleton/internal/middleware"
	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/jwtauth"
	"github.com/MQEnergy/go-skeleton/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func InitCommonGroup(r fiber.Router, handles ...fiber.Handler) {
	router := r.Group("/", handles...)
	{
		j := jwtauth.New(&vars.Config)

		router.Get("/", func(c *fiber.Ctx) error {
			return response.SuccessJSON(c, "", "")
		})
		router.Get("/ping", func(c *fiber.Ctx) error {
			return response.SuccessJSON(c, "", "pong")
		})

		// create jwt token
		router.Post("/token/create", func(c *fiber.Ctx) error {
			user := c.FormValue("user")
			pass := c.FormValue("pass")
			if user != "john" || pass != "doe" {
				return response.UnauthorizedException(c, "")
			}
			token, err := j.WithClaims(jwt.MapClaims{
				"name": user,
			}).GenerateToken()
			if err != nil {
				return response.UnauthorizedException(c, err.Error())
			}
			return response.SuccessJSON(c, "", token)
		})

		// parse jwt token
		router.Post("/token/view", middleware.AuthMiddleware(), func(ctx *fiber.Ctx) error {
			return response.SuccessJSON(ctx, "", fiber.Map{
				"uid":      ctx.GetRespHeader("uid"),
				"uuid":     ctx.GetRespHeader("uuid"),
				"role_ids": ctx.GetRespHeader("role_ids"),
			})
		})

		// 上传资源
		router.Post("/attachment/upload", backend.Attachment.Upload)

		// 登录
		router.Post("/backend/auth/login", backend.User.Login)

	}
}
