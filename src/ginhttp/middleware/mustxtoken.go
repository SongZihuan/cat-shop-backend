package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mustXTokenErrorData(debugMsgLst ...string) data.Data {
	debugMsg := ""

	if len(debugMsgLst) == 1 {
		debugMsg = debugMsgLst[0]
	}

	if debugMsg == "" {
		debugMsg = "Token过期"
	}

	return data.NewClientTokenExpireError(debugMsg)
}

func MustXTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value("user").(*model.User)
		if !ok {
			debugMsg, ok := c.Value(DebugTokenContextKey).(string)
			if !ok {
				debugMsg = "未知错误"
			}
			c.JSON(http.StatusOK, mustXTokenErrorData(debugMsg))
			return
		}

		if !user.CanLogin() {
			c.JSON(http.StatusOK, mustXTokenErrorData("用户状态不正确"))
			return
		}

		c.Next()
	}
}
