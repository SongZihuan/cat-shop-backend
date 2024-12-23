package addbag

import (
	"errors"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeWBagNotFound data.CodeType = -1
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

	if query.WuPinID <= 0 {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeWBagNotFound, "购物车未找到", "wupinID应该大于0"))
		return
	}

	bag, err := action.GetBagByWupinIDWithUser(user, query.WuPinID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeWBagNotFound, "购物车未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if !bag.WuPinShow || bag.WuPinID <= 0 || bag.WuPin == nil || !bag.WuPin.IsShow {
		c.JSON(http.StatusOK, data.NewNotSuccessData(CodeWBagNotFound, "购物车未找到"))
		return
	}

	_, err = action.AddBag(user, bag, query.Num)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccessData("添加成功"))
	return
}
