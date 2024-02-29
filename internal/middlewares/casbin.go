package middlewares

import (
	_ "embed"
	"strings"

	"go-skeleton/internal/vars"

	"go-skeleton/pkg/helper"
	"go-skeleton/pkg/response"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	casbin2 "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

//go:embed rbac_model.conf
var rbacModelConf string

// CasbinMiddleware casbin middleware
func CasbinMiddleware() fiber.Handler {
	adapter, _ := gormadapter.NewAdapterByDB(vars.DB)
	rc, _ := model.NewModelFromString(rbacModelConf)

	e, _ := casbin.NewEnforcer(rc, adapter)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	e.AddFunction("ParamsActMatch", ParamsActMatchFunc)
	_ = e.LoadPolicy()

	authz := casbin2.New(casbin2.Config{
		ModelFilePath: rbacModelConf,
		PolicyAdapter: adapter,
		Enforcer:      e,
		Lookup: func(c *fiber.Ctx) string {
			return c.GetRespHeader("uid")
		},
		Unauthorized: func(c *fiber.Ctx) error {
			if c.Path() == "/backend/auth/login" {
				return c.Next()
			}
			return response.UnauthorizedException(c, "unauthorized")
		}, // unauthorized handler
		Forbidden: func(c *fiber.Ctx) error {
			return response.ForbiddenException(c, "forbidden")
		}, // forbidden handler
	})
	return authz.RoutePermission()
}

// ParamsActMatchFunc 自定义规则函数
func ParamsActMatchFunc(args ...interface{}) (interface{}, error) {
	rAct := args[0].(string)
	pAct := args[1].(string)
	pActArr := strings.Split(pAct, ",")
	return helper.InAnySlice[string](pActArr, rAct), nil
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	key1 := strings.Split(name1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, name2), nil
}
