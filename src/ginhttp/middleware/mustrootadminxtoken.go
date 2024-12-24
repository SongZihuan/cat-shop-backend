package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"net/http"
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

func MustRootAdminXTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value(contextkey.UserKey).(*model.User)
		if !ok {
			c.JSON(http.StatusOK, mustRootAdminErrorData("用户获取失败"))
			return
		}

		if user.Type == modeltype.NormalUserType {
			c.JSON(http.StatusOK, mustAdminErrorData("普通用户权限不足"))
			return
		}

		if user.Type != modeltype.RootAdminUserType {
			c.JSON(http.StatusOK, mustAdminErrorData("普通管理员权限不足"))
			return
		}

		c.Next()
	}
}
