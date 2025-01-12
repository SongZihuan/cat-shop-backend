package updateuserinfo

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeNameError     data.CodeType = -1
	CodeWeChatError   data.CodeType = -2
	CodeEmailError    data.CodeType = -3
	CodeLocationError data.CodeType = -4
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.UserKey).(*model.User)
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

	err = useraction.UpdateUser(user, query.Name, query.Wechat, query.Email, query.Location)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
