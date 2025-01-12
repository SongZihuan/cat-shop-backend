package adminupdatexieyi

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	err = adminaction.AdminUpdateXieYi(query.Type, query.Content)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
