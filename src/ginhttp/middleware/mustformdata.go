package middleware

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const ContentTypeFormData = "multipart/form-data"
const ContentTypeEncoding = "charset=utf-8"

func MustFormData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			ct := c.GetHeader(header.RequestsContentType)
			cl := c.GetHeader(header.RequestsContentLength)

			if strings.Contains(ct, ContentTypeFormData) {
				abort.BadRequestsError(c, fmt.Errorf("Content-Type is not %s", ContentTypeFormData))
				return
			}

			if strings.Contains(ct, ContentTypeEncoding) {
				abort.BadRequestsError(c, fmt.Errorf("Content-Type encoding is not %s", ContentTypeEncoding))
				return
			}

			// Set max upload size (e.g., 10 MB)
			maxSize := int64(10 << 20) // 10 MB in bytes

			// Get the content length from the request headers
			contentLength, err := strconv.ParseInt(cl, 10, 64)
			if err != nil {
				abort.BadRequestsError(c, err)
				return
			}

			if contentLength > maxSize {
				abort.BadRequestsError(c, fmt.Errorf("Content-Length is too large: %d", contentLength))
				return
			}
		}

		c.Next()
	}
}
