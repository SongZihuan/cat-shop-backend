package adminpeoplepay

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	error2 "github.com/SongZihuan/cat-shop-backend/src/database/action/error"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
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

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	}

	record, err := adminaction.AdminGetBuyRecordByID(user, query.ID)
	if errors.Is(err, error2.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "交易非法", "未找到购物记录"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if record.IsBuyRecordCanNotPay() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotShort, "商户拒绝了此次交易", "商品不再出售"))
		return
	}

	err = useraction.SetBuyRecordPaySuccess(user, record)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("支付成功", "支付成功"))
	return
}
