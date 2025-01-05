package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mustAdminErrorData(debugMsgLst ...string) data.Data {
	debugMsg := ""

	if len(debugMsgLst) == 1 {
		debugMsg = debugMsgLst[0]
	}

	if debugMsg == "" {
		debugMsg = "权限不足"
	}

	return data.NewClientAdminError()
}

func MustAdminXTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value(contextkey.UserKey).(*model.User)
		if !ok {
			c.JSON(http.StatusOK, mustAdminErrorData("用户获取失败"))
			return
		}

		if user.Type == modeltype.NormalUserType {
			c.JSON(http.StatusOK, mustAdminErrorData("普通用户权限不足"))
			return
		}

		c.Next()
	}
}
