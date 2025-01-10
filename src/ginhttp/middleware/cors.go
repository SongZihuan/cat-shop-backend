package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/options"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Config().Yaml.Http.Cors.Enable() {
			if http.MethodOptions == c.Request.Method {
				if !options.Handler(c) {
					return
				}
			} else {
				if !cors.Handler(c) {
					return
				} else {
					c.Next()
				}
			}
		}
		c.Next()
	}
}
