package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestApiMiddleware() gin.HandlerFunc {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if config.Config().Yaml.Http.EnableTestAPI == "enable" {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			c.JSON(http.StatusOK, data.NewClientNotTestError())
		}
	}
}
