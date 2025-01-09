package adminimageupload

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
)

const Size3MB = 3 * 1024 * 1024
const Size2MB = 2 * 1024 * 1024

func Handler(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		abort.BadRequestsError(c, err)
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, NewError("非法请求"))
		return
	}

	var fileType modeltype.ImageType
	if query.Type == FileTypeXieyi {
		fileType = modeltype.XieYiImage
	} else if query.Type == FileTypeWupin {
		fileType = modeltype.WupinImage
	} else {
		c.JSON(http.StatusOK, NewError("文件类型错误"))
		return
	}

	alt, ok := modeltype.ImageAlt[fileType]
	if !ok {
		alt = ""
	}

	if query.File == nil {
		c.JSON(http.StatusOK, NewError("文件未上传"))
		return
	}

	if query.File.Size > Size2MB {
		c.JSON(http.StatusOK, NewError("文件太大"))
		return
	}

	file, err := query.File.Open()
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
		c.JSON(http.StatusOK, NewError("非图片"))
		return
	}

	img, errDB, errImg := adminaction.AdminNewImage(fileType, fileData)
	if errImg != nil || errDB != nil {
		c.JSON(http.StatusOK, NewError("系统错误"))
		return
	}

	c.JSON(http.StatusOK, NewSuccess(img.GetUrl(), alt))
}

func getImagePath() string {
	cfg := config.Config().Yaml.Http
	return cfg.BasePath + cfg.ResourcePath + modeltype.ImagePathV1
}
