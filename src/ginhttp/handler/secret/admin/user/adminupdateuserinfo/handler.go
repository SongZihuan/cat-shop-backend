package adminupdateuserinfo

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeNameError     data.CodeType = -3
	CodeWeChatError   data.CodeType = -4
	CodeEmailError    data.CodeType = -5
	CodeLocationError data.CodeType = -6
	CodeTypeError     data.CodeType = -7
	CodeStatusError   data.CodeType = -8
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.AdminUserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	self, ok := c.Value(contextkey.UserKey).(*model.User)
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
		c.JSON(http.StatusOK, data.NewCustomError(CodeEmailError, "邮箱过长"))
		return
	}

	if len(query.Email) != 0 && !utils.IsValidEmail(query.Email) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeEmailError, "邮箱非法"))
		return
	}

	if len(query.Location) > 160 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeLocationError, "地址过长"))
		return
	}

	errType, errStatus, errDB := action.AdminCreateUser(user, query.Name, query.Wechat, query.Email, query.Location, query.Status, query.Type, self.IsRootAdmin())
	if errType != nil {
		c.JSON(http.StatusOK, data.NewCustomError(CodeTypeError, "用户类型错误"))
		return
	} else if errStatus != nil {
		c.JSON(http.StatusOK, data.NewCustomError(CodeStatusError, "用户状态错误"))
		return
	} else if errDB != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("创建成功"))
}
