package getconfig

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	res, err := action.GetConfigLst()
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
	}

	c.JSON(http.StatusOK, NewJsonData(res))
	return
}
