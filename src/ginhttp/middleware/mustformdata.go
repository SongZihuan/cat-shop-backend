package middleware

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MustFormData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			ct := c.GetHeader(header.RequestsContentType)
			cl := c.GetHeader(header.RequestsContentLength)

			if ct != "multipart/form-data" {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			} else {
				// Set max upload size (e.g., 10 MB)
				maxSize := int64(10 << 20) // 10 MB in bytes

				// Get the content length from the request headers
				contentLength, err := strconv.ParseInt(cl, 10, 64)
				if err != nil {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}

				if contentLength > maxSize {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
			}
		}

		c.Next()
	}
}
