package adminupdatewupin

import (
	"errors"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
)

const Size3MB = 3 * 1024 * 1024
const Size2MB = 2 * 1024 * 1024

const (
	CodeFileTooBig    data.CodeType = -3
	CodeNotImage      data.CodeType = -4
	CodeWipinNotFound data.CodeType = -5
	CodeBadName       data.CodeType = -6
	CodeRealPrice     data.CodeType = -7
	CodeBadShopInfo   data.CodeType = -8
	CodeClassNotFound data.CodeType = -9
)

func Handler(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	newpic := ""
	if query.File != nil {
		if query.File.Size > Size2MB {
			c.JSON(http.StatusOK, data.NewCustomError(CodeFileTooBig, "文件太大"))
			return
		}

		file, err := query.File.Open()
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		fileData, err := io.ReadAll(file)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		mimeTp := mimetype.Detect(fileData).String()
		if !strings.HasPrefix("image/", mimeTp) {
			c.JSON(http.StatusOK, data.NewCustomError(CodeNotImage, "非图片"))
			return
		}

		img, errDB, errImg := action.NewImage(modeltype.AvatarImage, fileData)
		if errImg != nil {
			c.JSON(http.StatusOK, data.NewSystemUnknownError(errImg))
			return
		} else if errDB != nil {
			c.JSON(http.StatusOK, data.NewSystemDataBaseError(errDB))
			return
		}

		newpic = img.GetUrl()
	}

	if query.ID <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWipinNotFound, "商品未找到", "ID不应该小于0"))
		return
	}

	if len(query.Name) <= 0 || len(query.Name) > 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadName, "名称应在1-15个字符"))
		return
	}

	if query.RealPrice <= 0 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeRealPrice, "错误的真实售价"))
		return
	}

	if len(query.Ren) <= 0 || len(query.Ren) > 15 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadShopInfo, "商户名称应在1-15个字符"))
		return
	}

	if utils.InvalidPhone(query.Phone) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadShopInfo, "商户手机号错误"))
		return
	}

	if len(query.Location) < 10 || len(query.Location) > 150 {
		c.JSON(http.StatusOK, data.NewCustomError(CodeBadShopInfo, "商户定制应在10-150个字符"))
		return
	}

	cls, err := action.AdminGetClass(query.ClassID)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeClassNotFound, "类型未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	err = action.AdminUpdateWupin(query.ID, query.Name, newpic, cls, query.Tag, query.HotPrice.ToPriceNull(), query.RealPrice, query.Info, query.Ren, query.Phone, query.Email, query.Wechat, query.Location, query.Hot, query.Down)
	if errors.Is(err, action.ErrNotFound) {
		c.JSON(http.StatusOK, data.NewCustomError(CodeWipinNotFound, "商品未找到", "商品或分类未找到"))
		return
	} else if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(err))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
