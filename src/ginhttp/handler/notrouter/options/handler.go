package options

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/loadpath"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandlerAPI(c *gin.Context) bool {
	if config.Config().Yaml.Http.Cors.Disable() {
		return true
	}

	if !cors.Handler(c) {
		return false
	}

	c.Status(http.StatusNoContent)
	return true
}

func HandlerResource(c *gin.Context) bool {
	if config.Config().Yaml.Http.Cors.Disable() {
		return true
	}

	if !cors.Handler(c) {
		return false
	}

	c.Status(http.StatusNoContent)
	return true
}

func Handler(c *gin.Context) bool {
	api := loadpath.GetAPIPath()
	resource := loadpath.GetResourcePath()

	rawpath := utils.ProcessPath(c.Request.URL.Path)

	if api != "" && strings.HasPrefix(rawpath, api) {
		return HandlerAPI(c)
	} else if resource != "" && strings.HasPrefix(rawpath, resource) {
		return HandlerResource(c)
	} else {
		// 允许使用abort
		c.AbortWithStatus(http.StatusNotFound)
		return false
	}
}
