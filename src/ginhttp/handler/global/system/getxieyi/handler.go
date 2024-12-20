package getxieyi

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	db := database.DB()

	var res model.Xieyi
	err := db.Model(&model.Xieyi{}).Limit(1).Order("created_at desc").FirstOrCreate(&res, model.Xieyi{Data: ""}).Error
	if err != nil {
		c.JSON(200, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(200, NewJsonData(res.Data))
}
