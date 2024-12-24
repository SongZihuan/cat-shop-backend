package getsearch

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

const MaxPageSize = 20

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.PageSize > MaxPageSize || query.PageSize <= 0 {
		query.PageSize = MaxPageSize
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	res, err := action.GetSearchListWithShow(query.Search, query.Select, query.Page, query.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	maxcount, err := action.GetSearchCountWithShow(query.Search, query.Select)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, NewJsonData(res, maxcount))
	return
}
