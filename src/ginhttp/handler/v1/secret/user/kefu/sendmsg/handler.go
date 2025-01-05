package sendmsg

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeMsgMustBeNotEmpty     = -1
	CodeMsgLenMustLessThan160 = -2
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.UserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if len(query.Msg) == 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeMsgMustBeNotEmpty, "消息不能为空"))
		return
	}

	if len(query.Msg) > 160 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeMsgLenMustLessThan160, "消息不能超过160个字符"))
		return
	}

	err = action.SendMsgByUser(query.Msg, user)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("留言发送成功"))
	return
}
