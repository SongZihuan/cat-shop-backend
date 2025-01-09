package middleware

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"strings"
)

const AcceptJson = "application/json"
const AcceptEncoding = "charset=utf-8"

func MustAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptHeader := c.GetHeader(header.RequestsAccept)
		if !strings.Contains(acceptHeader, AcceptJson) {
			abort.NotAcceptError(c, AcceptJson)
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
			c.Writer.Header().Set(header.RequestsContentType, fmt.Sprintf("%s, %s", AcceptJson, AcceptEncoding))
		}
	}
}
