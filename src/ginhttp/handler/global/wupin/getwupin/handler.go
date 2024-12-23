package getwupin

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&Query{})
	if err != nil {
		c.JSON(http.StatusOK, NewMsgError(WupinNotFound, "未找到商品"))
		return
	}

	wupin, err := action.GetWupinByIDWithShow(query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, NewMsgError(WupinNotFound, "未找到商品"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(wupin))
	return
}
