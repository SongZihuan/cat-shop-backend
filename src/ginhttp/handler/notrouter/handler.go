package notrouter

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/options"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler404(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodOptions:
		options.Handler(c)
	default:
		c.JSON(http.StatusNotFound, data.NewClientNotFoundError())
	}
}

func Handler403(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodOptions:
		options.Handler(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, data.NewClientNotFoundError())
	}
}
