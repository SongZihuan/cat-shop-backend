package admingetclasslst

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

const MaxPageSize = 20

func Handler(c *gin.Context) {
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

	res, err := adminaction.AdminGetClassListByPage(query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	maxcount, err := adminaction.AdminGetClassCount()
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res, maxcount))
}
