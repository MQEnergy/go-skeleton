package jwtauth

import (
	"errors"
	"fmt"
	"time"

	"go-skeleton/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Info   interface{} `json:"info"`
	claims jwt.MapClaims
	token  *jwt.Token
	cfg    *config.Config
}

func New(cfg *config.Config) *JWT {
	claims := jwt.MapClaims{
		"iss": cfg.GetString("jwt.issuer"),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Second * cfg.GetDuration("jwt.expire"))),
	}
	return &JWT{
		claims: claims,
		cfg:    cfg,
	}
}

func (j *JWT) WithClaims(sub jwt.MapClaims) *JWT {
	j.claims["sub"] = sub
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	j.token = token
	return j
}

// GenerateToken encoded token
func (j *JWT) GenerateToken() (string, error) {
	return j.token.SignedString([]byte(j.cfg.GetString("jwt.secret")))
}

// ParseToken parse token
func (j *JWT) ParseToken(token string) (jwt.MapClaims, error) {
	_token, err := jwt.ParseWithClaims(token, j.claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.cfg.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, errors.New("err: " + err.Error())
	}
	if claims, ok := _token.Claims.(jwt.MapClaims); ok && _token.Valid {
		return claims, nil
	}
	return nil, errors.New("jwt 解析验证后失败")
}
