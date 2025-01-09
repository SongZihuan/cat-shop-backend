package addbag

import (
	"errors"
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
	CodeWBagNotFound  data.CodeType = -1
	CodeWupinNotFound data.CodeType = -2
	CodeBagDown       data.CodeType = -3
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

	if query.WupinID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "商品已下架", "wupinID应该大于0"))
		return
	}

	wupin, err := useraction.GetWupinByID(query.WupinID)
	if errors.Is(err, error2.ErrNotFound) || wupin == nil || wupin.IsWupinDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWupinNotFound, "商品以下架"))
		return
	}

	bag, err := useraction.GetUserWupinBag(user, wupin)
	if errors.Is(err, error2.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWBagNotFound, "购物车未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	} else if bag.IsBagDown() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBagDown, "购物车未找到"))
		return
	}

	_, err = useraction.AddBag(user, bag, query.Num)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("添加成功"))
	return
}
