package adminupdateclass

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeClassCanNotUpdate data.CodeType = -3
	CodeBadName           data.CodeType = -4
	CodeClassNotFound     data.CodeType = -5
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.ID <= 0 || query.ID == modeltype.ClassEmptyID {
		c.JSON(http.StatusOK, data.NewCustomError(CodeClassCanNotUpdate, "该类型不可编辑"))
		return
	}

	if len(query.Name) <= 0 || len(query.Name) > 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadName, "名称应在1-15个字符"))
		return
	}

	err = action.AdminUpdateClass(query.ID, query.Name, query.Show, query.Down)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeClassNotFound, "类型未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("添加成功"))
}
