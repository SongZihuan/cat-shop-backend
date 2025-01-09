package getwupin

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeWupinNotFound data.CodeType = -1
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "未找到商品", "ID应该大于0"))
		return
	}

	wupin, err := useraction.GetWupinByID(query.ID)
	if errors.Is(err, error2.ErrNotFound) || wupin == nil || wupin.IsWupinDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "未找到商品"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(wupin))
	return
}
