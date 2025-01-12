package get

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/loadpath"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandlerAPI403(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, data.NewClientNotFoundError())
}

func HandlerResource403(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
	_, _ = c.Writer.Write([]byte("Method Not Allowed"))
}

func HandlerAPI404(c *gin.Context) {
	c.JSON(http.StatusNotFound, data.NewClientNotFoundError())
}

func HandlerResource404(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
	_, _ = c.Writer.Write([]byte("Not Found"))
}

func Handler404(c *gin.Context) {
	api := loadpath.GetAPIPath()
	resource := loadpath.GetResourcePath()

	rawpath := utils.ProcessPath(c.Request.URL.Path)

	if strings.HasPrefix(rawpath, api) {
		HandlerAPI404(c)
	} else if strings.HasPrefix(rawpath, resource) {
		HandlerResource404(c)
	} else {
		// 允许使用abort
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func Handler403(c *gin.Context) {
	api := loadpath.GetAPIPath()
	resource := loadpath.GetResourcePath()

	rawpath := utils.ProcessPath(c.Request.URL.Path)

	if api != "" && strings.HasPrefix(rawpath, api) {
		HandlerAPI403(c)
	} else if resource != "" && strings.HasPrefix(rawpath, resource) {
		HandlerResource403(c)
	} else {
		// 允许使用abort
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}
