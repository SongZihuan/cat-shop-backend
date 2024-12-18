package middleware

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/jwttoken"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	RequestXTokenHeader = "X-Token"
)

const (
	ResponseXTokenHeader = "X-Token"
)

const (
	DebugTokenContextKey = "Debug:Token"
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
	token := c.GetHeader(RequestXTokenHeader)
	if token != "" {
		d, err := jwttoken.ParserUserToken(token)
		if err != nil {
			c.Set(DebugTokenContextKey, "Token解析失败: "+err.Error())
			return TokenStatusExpired
		}

		user, err := action.GetUserByID(d.Userid())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Set(DebugTokenContextKey, "用户未找到")
			return TokenStatusUserNotFound
		} else if err != nil {
			c.Set(DebugTokenContextKey, "Token解析失败: "+err.Error())
			return TokenStatusUserNotFound
		} else if !user.CanLogin() {
			c.Set(DebugTokenContextKey, "用户非正常状态")
			return TokenStatusUserNotOk
		}

		newToken := ""

		if d.IsNowReset() {
			newToken, err = jwttoken.CreateUserToken(user)
			if err == nil {
				c.Header(ResponseXTokenHeader, newToken)
			}
		}

		c.Set("OldToken", token)
		c.Set("NewToken", newToken)
		if newToken != "" {
			c.Set("Token", newToken)
		} else {
			c.Set("Token", token)
		}
		c.Set("UserID", user.ID)
		c.Set("User", user)
		return TokenStatusHasUser
	} else {
		c.Set("OldToken", "")
		c.Set("NewToken", "")
		c.Set("Token", "")
		c.Set("UserID", 0)
		c.Set("User", nil)
		c.Set(DebugTokenContextKey, "没有Token")
		return TokenStatusNotToken
	}
}

func XTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = HandlerToken(c)
		c.Next()
	}
}
