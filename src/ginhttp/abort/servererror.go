package abort

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerError(c *gin.Context, serverErr ...interface{}) {
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

	if len(serverErr) > 1 {
		serverErr = serverErr[:1]
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, data.NewSystemUnknownError(serverErr...))
}
