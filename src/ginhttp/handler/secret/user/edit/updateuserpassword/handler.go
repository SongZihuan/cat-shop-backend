package updateuserpassword

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodePasswordError data.CodeType = -1
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

	if !user.PasswordCheck(query.OldPassword) {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodePasswordError, "旧密码错误"))
		return
	}

	err = action.UpdateUserPassword(user, query.OldPassword, query.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccessData("更新成功"))
}
