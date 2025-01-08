package jwttoken

import (
	"errors"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
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
	aud := []string{user.GetLongName()}
	iat := jwt.NewNumericDate(now)
	iss := config.Config().Yaml.Jwt.Issuer
	nbf := jwt.NewNumericDate(now.Add(-1 * time.Second))
	sub := fmt.Sprintf("%sUSERTOKEN", strings.ToUpper(iss))
	exp := jwt.NewNumericDate(now.Add(time.Duration(config.Config().Yaml.Jwt.Hour) * time.Hour))
	rsm := time.Minute * time.Duration(config.Config().Yaml.Jwt.ResetMin)
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

func ParserUserToken(tokenString string) (data *Data, err error) {
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
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.userid <= 0 {
		return nil, errors.New("invalid token")
	}

	return &claims.Data, nil
}
