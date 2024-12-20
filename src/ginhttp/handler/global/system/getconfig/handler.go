package getconfig

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	db := database.DB()

	var res []model.Config
	err := db.Model(&model.Config{}).Where("key in ?", model.ConfigKey).Limit(len(model.ConfigKey)).Find(&res).Error
	if err != nil {
		c.JSON(200, data.NewSystemDataBaseError(err))
	}

	c.JSON(200, NewJsonData(res))
	return
}
