package repay

import (
	"errors"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

const (
	CodeBadRedirectTo      data.CodeType = -1
	CodeWupinNotFound      data.CodeType = -2
	CodeBuyRecordNotFound  data.CodeType = -3
	CodeWupinNotShort      data.CodeType = -4
	CodeBuyRecordStatusBad data.CodeType = -5
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

	record, err := useraction.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, error2.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if record.IsBuyRecordCanNotRepay() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordStatusBad, "商品不能再支付", "购物状态错误"))
		return
	}

	err = useraction.NewRepayRecord(user, record)
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
