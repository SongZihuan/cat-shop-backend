package ginplus

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Writer() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := writer.GinContextUseNewWriter(c)
		c.Next()
		_, err := w.WriteToHttp()
		if err != nil {
			// 允许使用c.Abort系列函数的地方
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				abort.ServerError(c, err)
				debugPrint("[ERROR] Recover: %v", err)
				if config.Config().Yaml.GlobalConfig.IsDebug() {
					panic(err)
				}
			}
		}()

		c.Next()
	}
}
