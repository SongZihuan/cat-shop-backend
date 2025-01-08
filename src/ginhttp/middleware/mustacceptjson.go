package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MustAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptHeader := c.GetHeader(header.RequestsAccept)
		if !strings.Contains(acceptHeader, "application/json") {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "Accept header must include application/json"})
			return
		}
		c.Next()
	}
}

func ReturnContentJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		ct := c.Writer.Header().Get(header.RequestsContentType)
		if ct != "" {
			c.Writer.Header().Set(header.RequestsContentType, "application/json; charset=utf-8")
		}
	}
}
