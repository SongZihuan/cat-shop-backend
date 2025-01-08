package middleware

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/SongZihuan/cat-shop-backend/src/jwttoken"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenStatus int

const (
	TokenStatusNotToken     TokenStatus = 1
	TokenStatusHasUser      TokenStatus = 2
	TokenStatusExpired      TokenStatus = 3
	TokenStatusUserNotFound TokenStatus = 4
	TokenStatusUserNotOk    TokenStatus = 5
)

func handlerToken(c *gin.Context) (*jwttoken.Data, *model.User, TokenStatus) {
	token := c.GetHeader(header.RequestXTokenHeader)
	if token != "" {
		tokenData, err := jwttoken.ParserUserToken(token)
		if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return nil, nil, TokenStatusExpired
		}

		user, err := action.MiddlewareGetUserByID(tokenData.Userid())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Set(contextkey.DebugTokenKey, "用户未找到")
			return nil, nil, TokenStatusUserNotFound
		} else if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return nil, nil, TokenStatusUserNotFound
		} else if !user.CanLogin() {
			c.Set(contextkey.DebugTokenKey, "用户非正常状态")
			return nil, nil, TokenStatusUserNotOk
		}

		c.Set(contextkey.TokenKey, token)
		c.Set(contextkey.UserIDKey, user.ID)
		c.Set(contextkey.UserKey, user)
		c.Set(contextkey.DebugTokenKey, "正常")
		return &tokenData, user, TokenStatusHasUser
	} else {
		c.Set(contextkey.DebugTokenKey, "没有Token")
		return nil, nil, TokenStatusNotToken
	}
}

func handlerResetToken(c *gin.Context, tokenData *jwttoken.Data, user *model.User, status TokenStatus) {
	if status == TokenStatusHasUser && httpStatusCheck(c) && user != nil && tokenData != nil && tokenData.IsNowReset() {
		if newToken, err := jwttoken.CreateUserToken(user); err == nil {
			c.Header(header.ResponseXTokenHeader, newToken)
		}
	}
}

func XTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenData, user, status := handlerToken(c)
		c.Next()
		handlerResetToken(c, tokenData, user, status)
	}
}
