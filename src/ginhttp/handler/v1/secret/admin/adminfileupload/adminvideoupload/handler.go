package adminvideoupload

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
)

const Size12MB = 12 * 1024 * 1024
const Size10MB = 10 * 1024 * 1024

func Handler(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(Size12MB); err != nil { // 32MB限制
		abort.BadRequestsError(c, err)
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, NewError("非法请求"))
		return
	}

	var fileType modeltype.VideoType
	if query.Type == FileTypeXieyi {
		fileType = modeltype.XieYiVideo
	} else if query.Type == FileTypeWupin {
		fileType = modeltype.WupinVideo
	} else {
		c.JSON(http.StatusOK, NewError("文件类型错误"))
		return
	}

	if query.File == nil {
		c.JSON(http.StatusOK, NewError("文件未上传"))
		return
	}

	if query.File.Size > Size10MB {
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
	if !strings.HasPrefix("video/", mimeTp) {
		c.JSON(http.StatusOK, NewError("非视频"))
		return
	}

	vid, errDB, errVid := action.AdminNewVideo(fileType, fileData)
	if errVid != nil || errDB != nil {
		c.JSON(http.StatusOK, NewError("系统错误"))
		return
	}

	c.JSON(http.StatusOK, NewSuccess(vid.GetUrl(), ""))
}
