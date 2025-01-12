package admingetclass

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CoddClassNotFound data.CodeType = -3
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CoddClassNotFound, "类型不存在", "类型ID不得小于等于0"))
		return
	} else if query.ID == modeltype.ClassEmptyID {
		c.JSON(http.StatusOK, NewClassEmptyJsonData())
		return
	}

	res, err := adminaction.AdminGetClass(query.ID)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res))
}
