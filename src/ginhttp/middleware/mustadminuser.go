package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mustAdminUserErrorData(debugMsgLst ...string) data.Data {
	debugMsg := ""

	if len(debugMsgLst) == 1 {
		debugMsg = debugMsgLst[0]
	}

	if debugMsg == "" {
		debugMsg = "权限不足"
	}

	return data.NewClientAdminUserNotFound()
}

func MustAdminUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Value(contextkey.AdminUserKey).(*model.User)
		if !ok {
			c.JSON(http.StatusOK, mustAdminUserErrorData("用户获取失败"))
			return
		}

		c.Next()
	}
}
