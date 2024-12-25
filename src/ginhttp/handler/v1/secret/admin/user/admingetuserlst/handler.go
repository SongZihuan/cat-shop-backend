package admingetuserlst

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const MaxPageSize = 20

func Handler(c *gin.Context) {
	self, ok := c.Value(contextkey.UserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("admin self not found"))
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

	res, err := action.GetUserByPage(query.Page, query.PageSize, self.IsRootAdmin())
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	maxcount, err := action.GetUserCount(self.IsRootAdmin())
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res, maxcount))
	return
}
