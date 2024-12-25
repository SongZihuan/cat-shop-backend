package adminadduser

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodePhoneError data.CodeType = -3
	CodeUserExists data.CodeType = -4
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(err))
		return
	}

	if !utils.IsChinaMainlandPhone(query.Phone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneError, "手机号不正确"))
		return
	}

	_, err = action.GetUserByPhone(query.Phone, false)
	if errors.Is(err, action.ErrNotFound) {
		_, err = action.CreateUser(query.Phone, query.Password)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else {
		c.JSON(http.StatusOK, data.NewCustomError(CodeUserExists, "该手机号已注册，请直接登录"))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("注册成功"))
	return
}
