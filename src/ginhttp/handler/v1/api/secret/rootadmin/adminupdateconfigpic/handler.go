package adminupdateconfigpic

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

const Size3MB = 3 * 1024 * 1024
const Size2MB = 2 * 1024 * 1024

const (
	CodeBadKey     = -3
	CodeNotFile    = -4
	CodeFileTooBig = -5
	CodeNotImage   = -6
	CodeNotPic     = -7
)

func Handler(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		abort.BadRequestsError(c, err)
		return
	}

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

	if query.Value == nil {
		canDelete, ok := modeltype.KeyCanDelete[keyType]
		if !ok {
			c.JSON(http.StatusOK, data.NewSystemUnknownError("无法获取配置项类型"))
			return
		}

		if canDelete {
			err := adminaction.AdminDeleteConfig(query.Key)
			if err != nil {
				c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
				return
			}
		} else {
			c.JSON(http.StatusOK, data.NewCustomError(CodeNotFile, "请上传文件"))
			return
		}
	}

	isPic, ok := modeltype.KeyIsPic[keyType]
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("无法获取配置项类型"))
		return
	}

	if !isPic {
		c.JSON(http.StatusOK, data.NewCustomError(CodeNotPic, "配置项为字符类型"))
		return
	}

	if query.Value.Size > Size2MB {
		c.JSON(http.StatusOK, data.NewCustomError(CodeFileTooBig, "文件太大"))
		return
	}

	file, err := query.Value.Open()
	if err != nil {
		abort.BadRequestsError(c, err)
		return
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		abort.BadRequestsError(c, err)
		return
	}

	mimeTp := mimetype.Detect(fileData).String()
	if !strings.HasPrefix("image/", mimeTp) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeNotImage, "非图片"))
		return
	}

	img, errDB, errImg := adminaction.AdminNewImage(modeltype.AvatarImage, fileData)
	if errImg != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(errImg))
		return
	} else if errDB != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errDB))
		return
	}

	err = adminaction.AdminUpdateConfigPic(query.Key, img)
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
