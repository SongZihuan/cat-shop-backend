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

func handlerToken(c *gin.Context) (*jwttoken.Data, TokenStatus) {
	token := c.GetHeader(header.RequestXTokenHeader)
	if token != "" {
		d, err := jwttoken.ParserUserToken(token)
		if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return nil, TokenStatusExpired
		}

		user, err := action.MiddlewareGetUserByID(d.Userid())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Set(contextkey.DebugTokenKey, "用户未找到")
			return nil, TokenStatusUserNotFound
		} else if err != nil {
			c.Set(contextkey.DebugTokenKey, "Token解析失败: "+err.Error())
			return nil, TokenStatusUserNotFound
		} else if !user.CanLogin() {
			c.Set(contextkey.DebugTokenKey, "用户非正常状态")
			return nil, TokenStatusUserNotOk
		}

		c.Set(contextkey.TokenKey, token)
		c.Set(contextkey.UserIDKey, user.ID)
		c.Set(contextkey.UserKey, user)
		c.Set(contextkey.DebugTokenKey, "正常")
		return nil, TokenStatusHasUser
	} else {
		c.Set(contextkey.DebugTokenKey, "没有Token")
		return nil, TokenStatusNotToken
	}
}

func handlerResetToken(c *gin.Context, token *jwttoken.Data, status TokenStatus) {
	if status == TokenStatusHasUser && token != nil && token.IsNowReset() && httpStatusCheck(c) {
		if userInterface, ok := c.Get(contextkey.UserKey); ok {
			if user, ok := userInterface.(*model.User); ok {
				newToken, err := jwttoken.CreateUserToken(user)
				if err == nil {
					c.Header(header.ResponseXTokenHeader, newToken)
				}
			}
		}
	}
}

func XTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, status := handlerToken(c)
		c.Next()
		handlerResetToken(c, token, status)
	}
}
