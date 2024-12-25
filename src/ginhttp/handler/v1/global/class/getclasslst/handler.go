package getclasslst

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ClassListLimit = 100

func Handler(c *gin.Context) {
	db := database.DB()

	var list = make([]model.Class, 0, 100)
	err := db.Model(&model.Class{}).Where("show = true").Limit(ClassListLimit).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(list))
}
