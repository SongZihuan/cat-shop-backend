package repay

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

const (
	CodeBadRedirectTo     data.CodeType = -1
	CodeWupinNotFound     data.CodeType = -2
	CodeBuyRecordNotFound data.CodeType = -3
	CodeWupinNotShort     data.CodeType = -4
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

	if strings.HasPrefix(query.RedirectTo, "http://") || strings.HasPrefix(query.RedirectTo, "https://") {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(CodeBadRedirectTo, "系统错误", "错误的RedirectTo地址"))
		return
	}

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到", "ID必须大于0"))
		return
	}

	if query.Type != modeltype.AliPay && query.Type != modeltype.WeChatPay {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("错误的支付类型"))
		return
	}

	record, err := action.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if !record.WuPin.IsWupinDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotShort, "购买记录未找到", "商品不再出售"))
		return
	}

	if record.WuPinID <= 0 || record.WuPin == nil {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "购物车未找到"))
		return
	}

	err = action.NewRepayRecord(user, record)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	payUrl := record.GetRepayPayUrl(query.Type, query.RedirectTo)
	if payUrl == "" {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("支付失败"))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(payUrl))
}
