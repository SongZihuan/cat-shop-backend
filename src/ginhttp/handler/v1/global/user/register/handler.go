package register

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/jwttoken"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodePhoneError data.CodeType = -1
	CodeUserExists data.CodeType = -2
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

	var user *model.User
	_, err = action.GetUserByPhone(query.Phone, false)
	if errors.Is(err, action.ErrNotFound) {
		user, err = action.CreateUser(query.Phone, query.Password)
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

	token, err := jwttoken.CreateUserToken(user)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(token))
	return
}
