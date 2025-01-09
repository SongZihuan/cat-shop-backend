package adminadduser

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
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

	if !utils.InvalidPhone(query.Phone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneError, "手机号不正确"))
		return
	}

	u, err := adminaction.AdminGetUserByPhone(query.Phone)
	if errors.Is(err, error2.ErrNotFound) || u == nil || u.Status == modeltype.DeleteUserStatus {
		_, err = adminaction.AdminCreateUser(query.Phone, query.Password)
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
