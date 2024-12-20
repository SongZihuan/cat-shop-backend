package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MustAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptHeader := c.GetHeader("Accept")
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
		ct := c.Writer.Header().Get("Content-Type")
		if ct != "" {
			c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
	}
}
