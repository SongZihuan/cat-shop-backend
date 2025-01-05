package adminupdateuserphone

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeUserIsDelete data.CodeType = -3
	CodePhoneExist   data.CodeType = -4
	CodePhoneError   data.CodeType = -5
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.AdminUserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	if user.IsDeleteUser() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeUserIsDelete, "用户已经被删除")) // 已经删除是用户无法执行操作
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

	u, err := action.AdminGetUserByPhone(query.NewPhone)
	if errors.Is(err, action.ErrNotFound) || u == nil || u.Status == modeltype.DeleteUserStatus {
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
