package backend

import (
	"errors"
	"strings"

	"github.com/MQEnergy/go-skeleton/internal/app/dao"
	"github.com/MQEnergy/go-skeleton/internal/app/model"

	"github.com/MQEnergy/go-skeleton/internal/request/user"
	"github.com/MQEnergy/go-skeleton/internal/vars"
	"github.com/MQEnergy/go-skeleton/pkg/helper"
	"github.com/MQEnergy/go-skeleton/pkg/jwtauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

var Auth = &AuthService{}

// Login
// @Description: 登录
// @receiver s
// @param reqParams
// @return interface{}
// @return error
// @author cx
func (s *AuthService) Login(reqParams *user.LoginReq) (fiber.Map, error) {
	var (
		isSuper   = 0 // 是否超级管理员 1：是 0：不是
		u         = dao.YfoAdmin
		err       error
		adminInfo *model.YfoAdmin
	)
	adminInfo, err = u.GetByAccount(reqParams.Account)
	if err != nil {
		return nil, errors.New("账号或密码不正确")
	}
	if adminInfo.Status != 1 {
		return nil, errors.New("用户已锁定，无法登录")
	}
	if adminInfo.Password != helper.GeneratePasswordHash(reqParams.Password, adminInfo.Salt) {
		return nil, errors.New("账号或密码不正确")
	}
	token, err := jwtauth.New(vars.Config).WithClaims(jwt.MapClaims{
		"id":       adminInfo.ID,
		"uuid":     adminInfo.UUID,
		"role_ids": adminInfo.RoleIds,
	}).GenerateToken()
	if err != nil {
		return nil, errors.New("登录失败")
	}
	if helper.InAnySlice(strings.Split(adminInfo.RoleIds, ","), "1") {
		isSuper = 1
	}
	return fiber.Map{
		"token": token,
		"info": fiber.Map{
			"id":       adminInfo.ID,
			"uuid":     adminInfo.UUID,
			"account":  adminInfo.Account,
			"avatar":   adminInfo.Avatar,
			"role_ids": adminInfo.RoleIds,
			"is_super": isSuper,
		},
	}, nil
}
