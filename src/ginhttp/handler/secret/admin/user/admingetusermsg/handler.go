package admingetusermsg

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const MaxPageSize = 20

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.AdminUserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.PageSize > MaxPageSize || query.PageSize <= 0 {
		query.PageSize = MaxPageSize
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	res, err := action.GetMsgByPageAndUser(user, query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	maxcount, err := action.GetMsgCountWithUser(user)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res, maxcount))
	return
}
