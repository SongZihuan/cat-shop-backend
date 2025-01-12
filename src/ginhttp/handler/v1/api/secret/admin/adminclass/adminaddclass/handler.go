package adminaddclass

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

const (
	CodeBadName data.CodeType = -3
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if len(query.Name) <= 0 || len(query.Name) > 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadName, "名称应在1-15个字符"))
		return
	}

	err = adminaction.AdminAddClass(query.Name, query.Show, query.Down)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("添加成功"))
}
