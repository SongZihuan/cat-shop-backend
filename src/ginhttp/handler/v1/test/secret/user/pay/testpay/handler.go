package testpay

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
	CodeBuyRecordNotFound  data.CodeType = -1
	CodeRepeatTransactions data.CodeType = -2
	CodePayFail            data.CodeType = -3
	CodeWupinNotShort      data.CodeType = -4
)

const DefaultFailRate int = 10

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
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	}

	if query.FailRate == -1 {
		query.FailRate = DefaultFailRate
	} else if query.FailRate < 0 {
		query.FailRate = 0
	} else if query.FailRate > 100 {
		query.FailRate = 100
	}

	record, err := action.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if record.IsBuyRecordCanNotPay() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotShort, "商户拒绝了此次交易", "商品不再出售"))
		return
	}

	if utils.Rand().Intn(100) < query.FailRate { // 10%概率支付失败
		err := action.SetBuyRecordPayFail(user, record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusOK, data.NewCustomError(CodePayFail, "支付失败", "支付失败"))
		return
	} else {
		err := action.SetBuyRecordPaySuccess(user, record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusOK, data.NewSuccess("支付成功", "支付成功"))
		return
	}
}
