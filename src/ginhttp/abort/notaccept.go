package abort

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotAcceptError(c *gin.Context, accept string) {
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

	c.AbortWithStatusJSON(http.StatusBadRequest, data.NewClientNotAccept(accept))
}
