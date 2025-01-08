package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database"
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllReady() gin.HandlerFunc {
	if flagparser.IsReady() && config.IsReady() && database.IsReady() {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}
