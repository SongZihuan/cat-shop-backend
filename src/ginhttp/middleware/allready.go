package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database"
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/logger"
	"github.com/gin-gonic/gin"
)

func AllReady() gin.HandlerFunc {
	return func(c *gin.Context) {
		if flagparser.IsReady() && config.IsReady() && logger.IsReady() && database.IsReady() {
			c.Next()
			return
		}

		abort.ServerError(c, "系统未准备好")
	}
}
