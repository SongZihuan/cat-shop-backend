package testpay

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeBuyRecordNotFound  data.CodeType = 1
	CodeRepeatTransactions data.CodeType = 2
	CodePayFail            data.CodeType = 3
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&Query{})
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if query.ID <= 0 {
		c.JSON(http.StatusNotFound, data.NewNotSuccessData(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	}

	record, err := action.GetBuyRecordByID(query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusNotFound, data.NewNotSuccessData(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	if record.Status != modeltype.WaitPayCheck && record.Status != modeltype.PayCheckFail {
		c.JSON(http.StatusNotFound, data.NewNotSuccessData(CodeRepeatTransactions, "重复交易", "购物记录状态不正确"))
		return
	}

	if utils.Rand().Intn(100) < 10 { // 10%概率支付失败
		err := action.SetBuyRecordPayFail(record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusNotFound, data.NewNotSuccessData(CodePayFail, "支付失败", "支付失败"))
		return
	} else {
		err := action.SetBuyRecordPaySuccess(record)
		if err != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
			return
		}

		c.JSON(http.StatusNotFound, data.NewSuccessData("支付成功", "支付成功"))
		return
	}
}
