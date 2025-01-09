package abort

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/writer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestsError(c *gin.Context, serverErr ...interface{}) {
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

	if len(serverErr) == 1 {
		if err, ok := serverErr[0].(error); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.NewClientBadRequests(err))
			return
		} else if msg, ok := serverErr[0].(string); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.NewClientBadRequests(fmt.Errorf("%s", msg)))
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, data.NewClientBadRequests(fmt.Errorf("请求错误")))
}
