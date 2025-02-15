package updateuseravtar

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/useraction"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/contextkey"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
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

const (
	CodeFileTooBig    data.CodeType = -3
	CodeNotImage      data.CodeType = -4
	CodeFileNotUpload data.CodeType = -5
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.UserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		abort.BadRequestsError(c, err)
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&query, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

	if query.File == nil {
		c.JSON(http.StatusOK, data.NewCustomError(CodeFileNotUpload, "文件未上传"))
		return
	}

	if query.File.Size > Size2MB {
		c.JSON(http.StatusOK, data.NewCustomError(CodeFileTooBig, "文件太大"))
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
		c.JSON(http.StatusOK, data.NewCustomError(CodeNotImage, "非图片"))
		return
	}

	img, errDB, errImg := useraction.NewImage(modeltype.AvatarImage, fileData)
	if errImg != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(errImg))
		return
	} else if errDB != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errDB))
		return
	}

	err = useraction.UpdateUserAvatar(user, img.GetUrl())
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errImg))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
