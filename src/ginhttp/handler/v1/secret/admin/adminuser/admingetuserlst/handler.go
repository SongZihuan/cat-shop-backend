package admingetuserlst

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/rootaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
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

	var res []model.User
	if self.IsRootAdmin() {
		res, err = rootaction.RootAdminGetUserByPage(query.Page, query.PageSize)
	} else {
		res, err = adminaction.AdminGetUserByPage(query.Page, query.PageSize)
	}
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	var maxcount int
	if self.IsRootAdmin() {
		maxcount, err = rootaction.RootAdminGetUserCount()
	} else {
		maxcount, err = adminaction.AdminGetUserCount()
	}
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res, maxcount))
	return
}
