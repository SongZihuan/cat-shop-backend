package getbaglstbypage

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/contextkey"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const DefaultLimit = 10
const MaxLimit = 20

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.UserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.Limit <= 0 {
		query.Limit = DefaultLimit
	}

	if query.Limit > MaxLimit {
		query.Limit = MaxLimit
	}

	if query.Offset < 0 {
		query.Offset = 0
	}

	res, err := action.GetBagListByUser(user, query.Limit, query.Offset, false)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res))
}
