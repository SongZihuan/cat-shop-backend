package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
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
		_, ok := c.Value("user").(*model.User)
		if !ok {
			debugMsg, ok := c.Value(DebugTokenContextKey).(string)
			if !ok {
				debugMsg = "未知错误"
			}
			c.JSON(http.StatusOK, mustXTokenErrorData(debugMsg))
			return
		}

		c.Next()
	}
}
