package abort

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResourceNotAccept(c *gin.Context, accept string) {
	defer func() {
		recover()
	}()

	w, ok := c.Writer.(*writer.Writer)
	if !ok {
		// 允许使用c.Abort系列函数的地方
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	err := w.Reset()
	if err != nil {
		// 允许使用c.Abort系列函数的地方
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	w.WriteString(fmt.Sprintf("406 Not Acceptable: %s", accept))
	c.AbortWithStatus(http.StatusNotAcceptable)
}
