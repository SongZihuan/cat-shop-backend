package testpay

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBuyRecordNotFound  data.CodeType = 1
	CodeRepeatTransactions data.CodeType = 2
	CodePayFail            data.CodeType = 3
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

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	}

	record, err := action.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	if record.Status != modeltype.WaitPayCheck && record.Status != modeltype.PayCheckFail {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeRepeatTransactions, "重复交易", "购物记录状态不正确"))
		return
	}

	if utils.Rand().Intn(100) < 10 { // 10%概率支付失败
		err := action.SetBuyRecordPayFail(record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusOK, data.NewNotSuccessData(CodePayFail, "支付失败", "支付失败"))
		return
	} else {
		err := action.SetBuyRecordPaySuccess(user, record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusOK, data.NewSuccessData("支付成功", "支付成功"))
		return
	}
}
