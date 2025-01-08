package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllowMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !header.IsAllowMethods(c.Request.Method) {
			c.JSON(http.StatusMethodNotAllowed, data.NewClientNotFoundError())
			return
		}

		c.Next()
	}
}
