package middleware

import (
	_ "embed"
	"strings"

	"github.com/MQEnergy/go-skeleton/internal/vars"

	"github.com/MQEnergy/go-skeleton/pkg/helper"
	"github.com/MQEnergy/go-skeleton/pkg/response"

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
	if vars.DB == nil {
		return func(c *fiber.Ctx) error {
			return nil
		}
	}
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
			return response.UnauthorizedException(c, "权限不足")
		}, // unauthorized handler
		Forbidden: func(c *fiber.Ctx) error {
			return response.ForbiddenException(c, "forbidden")
		}, // forbidden handler
	})
	return authz.RoutePermission()
}

// ParamsActMatchFunc 自定义规则函数 method
func ParamsActMatchFunc(args ...interface{}) (interface{}, error) {
	rAct := args[0].(string)
	pAct := args[1].(string)
	pActArr := strings.Split(pAct, ",")
	if len(pActArr) > 1 {
		return helper.InAnySlice[string](pActArr, rAct), nil
	}
	return pActArr[0] == rAct, nil
}

// ParamsMatchFunc 自定义规则函数 path
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	rObj := args[0].(string)
	pObj := args[1].(string)
	rObj1 := strings.Split(rObj, "?")[0]
	return util.KeyMatch2(rObj1, pObj), nil
}
