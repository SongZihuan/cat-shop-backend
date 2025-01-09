package options

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandlerAPI(c *gin.Context) {
	if config.Config().Yaml.Http.Cors.Enable() {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, data.NewClientCorsError("系统未启动Cors跨域模式"))
		return
	}

	if !cors.Handler(c) {
		return
	}

	c.Status(http.StatusNoContent)
}

func HandlerResource(c *gin.Context) {
	if config.Config().Yaml.Http.Cors.Enable() {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	if !cors.Handler(c) {
		return
	}

	c.Status(http.StatusNoContent)
}

func Handler(c *gin.Context) {
	base := config.Config().Yaml.Http.BasePath
	api := utils.ProcessPath(base + config.Config().Yaml.Http.ApiPath)
	resource := utils.ProcessPath(base + config.Config().Yaml.Http.ResourcePath)
	rawpath := utils.ProcessPath(c.Request.URL.Path)

	if strings.HasPrefix(rawpath, api) {
		HandlerAPI(c)
	} else if strings.HasPrefix(rawpath, resource) {
		HandlerResource(c)
	} else {
		// 允许使用abort
		c.AbortWithStatus(http.StatusNotFound)
	}
}
