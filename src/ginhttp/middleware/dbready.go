package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/database"
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
