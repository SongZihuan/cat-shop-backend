package login

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/jwttoken"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodePhoneError    data.CodeType = -1
	CodePasswordError data.CodeType = -2
	CodeUserFreeze    data.CodeType = -3
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

	user, err := action.GetUserByPhone(query.Phone)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePhoneError, "用户不存在或密码错误", "用户不存在"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	if !user.CanLogin() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeUserFreeze, "用户被冻结"))
		return
	}

	if !user.PasswordCheck(query.Password) {
		c.JSON(http.StatusOK, data.NewCustomError(CodePasswordError, "用户不存在或密码错误", "密码不匹配"))
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
