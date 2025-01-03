package middleware

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/header"
	"github.com/SuperH-0630/cat-shop-back/src/jwttoken"
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

func HandlerToken(c *gin.Context) TokenStatus {
	token := c.GetHeader(header.RequestXTokenHeader)
	if token != "" {
		d, err := jwttoken.ParserUserToken(token)
		if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return TokenStatusExpired
		}

		user, err := action.GetUserByID(d.Userid(), false)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Set(contextkey.DebugTokenKey, "用户未找到")
			return TokenStatusUserNotFound
		} else if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return TokenStatusUserNotFound
		} else if !user.CanLogin() {
			c.Set(contextkey.DebugTokenKey, "用户非正常状态")
			return TokenStatusUserNotOk
		}

		newToken := ""

		if d.IsNowReset() {
			newToken, err = jwttoken.CreateUserToken(user)
			if err == nil {
				c.Header(header.ResponseXTokenHeader, newToken)
			}
		}

		c.Set(contextkey.OldTokenKey, token)
		c.Set(contextkey.NewTokenKey, newToken)
		if newToken != "" {
			c.Set(contextkey.TokenKey, newToken)
		} else {
			c.Set(contextkey.TokenKey, token)
		}
		c.Set(contextkey.UserIDKey, user.ID)
		c.Set(contextkey.UserKey, user)
		c.Set(contextkey.DebugTokenKey, "正常")
		return TokenStatusHasUser
	} else {
		c.Set(contextkey.DebugTokenKey, "没有Token")
		return TokenStatusNotToken
	}
}

func XTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = HandlerToken(c)
		c.Next()
	}
}
