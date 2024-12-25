package testpay

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/handler/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBuyRecordNotFound  data.CodeType = -1
	CodeRepeatTransactions data.CodeType = -2
	CodePayFail            data.CodeType = -3
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
	}

	if record.Status != modeltype.WaitPayCheck && record.Status != modeltype.PayCheckFail {
		c.JSON(http.StatusOK, data.NewCustomError(CodeRepeatTransactions, "重复交易", "购物记录状态不正确"))
		return
	}

	if utils.Rand().Intn(100) < query.FailRate { // 10%概率支付失败
		err := action.SetBuyRecordPayFail(record)
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
