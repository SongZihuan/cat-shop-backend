package adminupdateuserphone

import (
	"errors"
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
	CodePasswordError data.CodeType = -3
	CodePhoneExist    data.CodeType = -4
	CodePhoneError    data.CodeType = -5
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

	if user.IsRootAdmin() && !self.IsRootAdmin() {
		c.JSON(http.StatusOK, data.NewClientAdminUserNoPermission())
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if !utils.IsChinaMainlandPhone(query.NewPhone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneError, "手机号不正确"))
		return
	}

	_, err = action.GetUserByPhone(query.NewPhone, false)
	if errors.Is(err, action.ErrNotFound) {
		err := action.AdminUpdateUserPhone(user, query.NewPhone)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneExist, "该手机号已注册，请直接登录"))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
