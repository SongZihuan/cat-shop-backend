package adminaddbag

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
	CodeBagNotFound data.CodeType = -3
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.AdminUserKey).(*model.User)
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

	if query.WupinID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到", "wupinID应该大于0"))
		return
	}

	wupin, err := action.AdminGetWupinByID(query.WupinID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到", "商品未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if wupin.IsWupinDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到", "物品下架"))
		return
	}

	bag, err := action.AdminGetBagByWupinIDWithUser(user, wupin)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if bag.IsBagDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagNotFound, "购物车未找到"))
		return
	}

	_, err = action.AdminAddBag(user, bag, query.Num)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("添加成功"))
	return
}
