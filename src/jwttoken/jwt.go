package jwttoken

import (
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

func CreateUserToken(user *model.User) (t string, err error) {
	if config.IsReady() {
		panic("config is not ready")
	}

	defer func() {
		_ = recover()
		err = errors.New("invalid token")
	}()

	now := time.Now()

	jti := fmt.Sprintf("%d%d%d", now.UnixNano(), user.ID, utils.Rand().Intn(100))
	aud := []string{user.GetName()}
	iat := jwt.NewNumericDate(now)
	iss := config.Config().Yaml.Jwt.Issuer
	nbf := jwt.NewNumericDate(now.Add(-1 * time.Second))
	sub := fmt.Sprintf("%sUSERTOKEN", strings.ToUpper(iss))
	exp := jwt.NewNumericDate(now.Add(config.Config().Yaml.Jwt.Hour * time.Hour))

	rsm := time.Minute * config.Config().Yaml.Jwt.ResetMin
	rst := exp.Add(-rsm)

	c := Claims{
		Data: Data{
			userid:    user.ID,
			resetmin:  rsm,
			resettime: rst,
		},

		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  aud,
			ExpiresAt: exp,
			ID:        jti,
			Issuer:    iss,
			IssuedAt:  iat,
			NotBefore: nbf,
			Subject:   sub,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(config.Config().Yaml.Jwt.Secret))
}

func ParserUserToken(tokenString string) (data Data, err error) {
	if config.IsReady() {
		panic("config is not ready")
	}

	defer func() {
		_ = recover()
		err = errors.New("invalid token")
	}()

	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config().Yaml.Jwt.Secret), nil
	})
	if err != nil {
		return Data{}, err
	} else if !token.Valid {
		return Data{}, errors.New("invalid token")
	}

	if claims.userid <= 0 {
		return Data{}, errors.New("invalid token")
	}

	return claims.Data, nil
}
