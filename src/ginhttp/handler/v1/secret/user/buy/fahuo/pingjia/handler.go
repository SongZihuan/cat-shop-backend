package pingjia

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBuyRecordNotFound data.CodeType = -1
	CodeStatusError       data.CodeType = -2
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

	record, err := action.GetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	err = action.BuyRecordPingJia(user, record, query.IsGood)
	if _, ok := action.IsBuyRecordStatusError(err); ok {
		c.JSON(http.StatusOK, data.NewCustomError(CodeStatusError, err.Error()))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("评价成功"))
}
