package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpStatusCheck(c *gin.Context) bool {
	return c.Writer.Status() == http.StatusOK || c.Writer.Status() == http.StatusNoContent
}
