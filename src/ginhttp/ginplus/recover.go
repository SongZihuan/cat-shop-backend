package ginplus

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				debugPrint("[ERROR] Recover: %v", err)
				if config.Config().Yaml.Global.IsDebug() {
					panic(err)
				}
			}
		}()

		c.Next()
	}
}
