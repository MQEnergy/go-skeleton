package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-skeleton/internal/middlewares"
	"go-skeleton/internal/vars"
	"go-skeleton/pkg/jwtauth"
	"go-skeleton/pkg/response"
)

func InitCommonGroup(r fiber.Router, middleware ...fiber.Handler) {
	router := r.Group("/", middleware...)
	{
		j := jwtauth.New(vars.Config)

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
			token, err := j.ApplySub(map[string]interface{}{
				"name": user,
			}).ApplyClaims().GenerateToken()
			if err != nil {
				return response.UnauthorizedException(c, err.Error())
			}
			return response.SuccessJSON(c, "", token)
		})

		// parse jwt token
		router.Post("/token/view", middlewares.AuthMiddleware(), func(ctx *fiber.Ctx) error {
			return response.SuccessJSON(ctx, "", fiber.Map{
				"uid":      ctx.GetRespHeader("uid"),
				"uuid":     ctx.GetRespHeader("uuid"),
				"role_ids": ctx.GetRespHeader("role_ids"),
			})
		})

	}
}
