package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DBReady() gin.HandlerFunc {
	if database.IsReady() {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}
