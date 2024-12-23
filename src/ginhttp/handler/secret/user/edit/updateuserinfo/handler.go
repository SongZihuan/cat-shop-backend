package updateuserinfo

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeNameError     data.CodeType = 1
	CodeWeChatError   data.CodeType = 2
	CodeEmailError    data.CodeType = 3
	CodeLocationError data.CodeType = 4
)

func Handler(c *gin.Context) {
	user, ok := c.Value("User").(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&Query{}, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if len(query.Name) == 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeNameError, "名字必须设定"))
		return
	}

	if len(query.Name) > 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeNameError, "名字过长"))
		return
	}

	if len(query.Wechat) > 45 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWeChatError, "微信过长"))
		return
	}

	if len(query.Email) > 45 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWeChatError, "邮箱过长"))
		return
	}

	if len(query.Email) != 0 && !utils.IsValidEmail(query.Email) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWeChatError, "邮箱非法"))
		return
	}

	if len(query.Location) > 160 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWeChatError, "地址过长"))
		return
	}

	err = action.UpdateUser(user, query.Name, query.Wechat, query.Email, query.Location)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
