package options

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	if config.Config().Yaml.Http.Cors.Enable() {
		c.JSON(http.StatusMethodNotAllowed, data.NewClientCorsError("系统未启动Cors跨域模式"))
		return
	}

	if !cors.Handler(c) {
		return
	}

	c.Status(http.StatusNoContent)
}
