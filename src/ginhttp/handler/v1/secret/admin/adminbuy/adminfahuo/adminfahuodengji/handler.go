package adminfahuodengji

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBuyRecordNotFound data.CodeType = -3
	CodeKuaiDiError       data.CodeType = -4
	CodeKuaiDiNumError    data.CodeType = -5
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

	if len(query.KuaiDi) <= 0 || len(query.KuaiDi) >= 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeKuaiDiError, "快递名称错误", "快递名称长度错误"))
		return
	}

	if len(query.KuaiDiNum) <= 0 || len(query.KuaiDiNum) >= 45 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeKuaiDiNumError, "快递单号错误", "快递单号长度错误"))
		return
	}

	record, err := action.AdminGetBuyRecordByIDAndUser(user, query.ID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBuyRecordNotFound, "购买记录未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	err = action.AdminBuyRecordFaHuoDengJi(user, record, query.KuaiDi, query.KuaiDiNum)
	if _, ok := action.IsBuyRecordStatusError(err); ok {
		c.JSON(http.StatusOK, data.NewCustomError(CodeStatusError, err.Error()))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("确认到货成功"))
}
