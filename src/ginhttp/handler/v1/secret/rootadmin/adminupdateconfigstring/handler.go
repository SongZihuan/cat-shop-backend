package adminupdateconfigstring

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeBadKey          = -3
	CodeKeyCanNotDelete = -4
	CodeKeyNotString    = -5
)

func Handler(c *gin.Context) {
	query := Query{}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	_, ok := modeltype.ConfigKeyMap[query.Key]
	if !ok {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadKey, "配置项错误"))
		return
	}

	keyType, ok := modeltype.ConfigType[query.Key]
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("无法获取配置项类型"))
		return
	}

	canDelete, ok := modeltype.KeyCanDelete[keyType]
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("无法获取配置项类型"))
		return
	}

	if query.Value == "" && !canDelete {
		c.JSON(http.StatusOK, data.NewCustomError(CodeKeyCanNotDelete, "配置项不可删除"))
		return
	}

	isPic, ok := modeltype.KeyIsPic[keyType]
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("无法获取配置项类型"))
		return
	}

	if isPic {
		c.JSON(http.StatusOK, data.NewCustomError(CodeKeyNotString, "配置项为图片类型"))
		return
	}

	err = adminaction.AdminUpdateConfigString(query.Key, query.Value)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
