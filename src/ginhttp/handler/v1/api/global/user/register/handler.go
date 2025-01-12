package register

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/jwttoken"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
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

	if !utils.InvalidPhone(query.Phone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneError, "手机号不正确"))
		return
	}

	var user *model.User
	_, err = useraction.GetUserByPhone(query.Phone)
	if errors.Is(err, error2.ErrNotFound) {
		user, err = adminaction.AdminCreateUser(query.Phone, query.Password)
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
