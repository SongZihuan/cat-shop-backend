package getwupin

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeWupinNotFound data.CodeType = -1
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&Query{})
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "未找到商品", "ID应该大于0"))
		return
	}

	wupin, err := action.GetWupinByIDWithShow(query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "未找到商品"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(wupin))
	return
}
