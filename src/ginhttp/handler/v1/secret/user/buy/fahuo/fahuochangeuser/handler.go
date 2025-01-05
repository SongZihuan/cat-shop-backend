package fahuochangeuser

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBuyRecordNotFound data.CodeType = -1
	CodeBadName           data.CodeType = -2
	CodeBadPhone          data.CodeType = -3
	CodeBadLocation       data.CodeType = -4
	CodeBadEmail          data.CodeType = -5
	CodeStatusError       data.CodeType = -6
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

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	}

	if len(query.UserName) <= 0 {
		query.UserName = user.Name
	} else if len(query.UserName) >= 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadName, "购买人姓名不对", "购买人姓名太长"))
		return
	}

	if len(query.UserPhone) <= 0 {
		query.UserPhone = user.Phone
	} else if !utils.IsChinaMainlandPhone(query.UserPhone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadPhone, "购买人联系电话不对"))
		return
	}

	if len(query.UserLocation) <= 0 || len(query.UserLocation) >= 160 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadLocation, "购买人联系地址不对"))
		return
	}

	if len(query.UserEmail) > 0 && !utils.IsValidEmail(query.UserEmail) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadEmail, "错误的邮件地址"))
		return
	}

	if len(query.UserRemark) > 160 {
		query.UserRemark = query.UserRemark[0:160]
	}

	record, err := action.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	err = action.BuyRecordChangeUser(user, record, query.UserName, query.UserPhone, query.UserLocation, query.UserWechat, query.UserEmail, query.UserRemark)
	if _, ok := action.IsBuyRecordStatusError(err); ok {
		c.JSON(http.StatusOK, data.NewCustomError(CodeStatusError, err.Error()))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("确认到货成功"))
}
