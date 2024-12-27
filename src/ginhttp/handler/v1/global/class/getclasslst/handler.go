package getclasslst

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ClassListLimit = 100

func Handler(c *gin.Context) {
	res, err := action.GetClassList(ClassListLimit, false, true, false)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res))
}
