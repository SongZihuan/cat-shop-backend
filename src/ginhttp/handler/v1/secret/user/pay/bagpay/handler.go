package bagpay

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

const (
	CodeBadRedirectTo data.CodeType = -1
	CodeBagNotFound   data.CodeType = -2
	CodeBadName       data.CodeType = -4
	CodeBadPhone      data.CodeType = -5
	CodeBadLocation   data.CodeType = -6
	CodeBadEmail      data.CodeType = -7
	CodeWupinNotFound data.CodeType = -9
	CodeWupinNotShort data.CodeType = -10
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

	if query.BagID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到", "bagID必须大于0"))
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

	if query.Type != modeltype.AliPay && query.Type != modeltype.WeChatPay {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("错误的支付类型"))
		return
	}

	bag, err := action.GetBagByIDAndUser(user, query.BagID, false)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if !bag.IsBagDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotShort, "购物车未找到", "商品不再出售"))
		return
	}

	if bag.WuPinID <= 0 || bag.WuPin == nil {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "购物车未找到"))
		return
	}

	record, err := action.NewBagBuyRecord(user, bag, query.UserName, query.UserPhone, query.UserLocation, query.UserWechat, query.UserEmail, query.UserRemark)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	payUrl := record.GetBagPayUrl(query.Type, query.RedirectTo)
	if payUrl == "" {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("支付失败"))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(payUrl))
}
