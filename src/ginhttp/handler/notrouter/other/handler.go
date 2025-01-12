package other

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler404(c *gin.Context) {
	Handler403(c)
}

func Handler403(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
	_, _ = c.Writer.Write([]byte("Method Not Allowed"))
}
