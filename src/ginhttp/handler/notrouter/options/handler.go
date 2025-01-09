package options

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/loadpath"
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
	api := loadpath.GetAPIPath()
	resource := loadpath.GetResourcePath()

	rawpath := utils.ProcessPath(c.Request.URL.Path)

	if api != "" && strings.HasPrefix(rawpath, api) {
		HandlerAPI(c)
	} else if resource != "" && strings.HasPrefix(rawpath, resource) {
		HandlerResource(c)
	} else {
		// 允许使用abort
		c.AbortWithStatus(http.StatusNotFound)
	}
}
