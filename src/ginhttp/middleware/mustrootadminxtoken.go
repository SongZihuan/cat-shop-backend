package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
)

func mustRootAdminErrorData(debugMsgLst ...string) data.Data {
	debugMsg := ""

	if len(debugMsgLst) == 1 {
		debugMsg = debugMsgLst[0]
	}

	if debugMsg == "" {
		debugMsg = "重要权限不足"
	}

	return data.NewClientRootAdminError()
}

func MustRotAdminXTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value("user").(*model.User)
		if !ok {
			c.JSON(200, mustRootAdminErrorData("用户获取失败"))
			return
		}

		if user.Type == model.NormalUserType {
			c.JSON(200, mustAdminErrorData("普通用户权限不足"))
			return
		}

		if user.Type == model.AdminUserType {
			c.JSON(200, mustAdminErrorData("普通管理员权限不足"))
			return
		}

		c.Next()
	}
}
