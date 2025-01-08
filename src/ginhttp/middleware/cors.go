package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if http.MethodOptions == c.Request.Method {
			if config.Config().Yaml.Http.Cors.Disable() {
				c.JSON(http.StatusMethodNotAllowed, data.NewClientCorsError("系统未启动Cors跨域模式"))
				return
			} else if !cors.Handler(c) {
				return
			} else {
				c.Status(http.StatusNoContent)
			}
		} else {
			if config.Config().Yaml.Http.Cors.Enable() && !cors.Handler(c) {
				return
			} else {
				c.Next()
			}
		}
	}
}
