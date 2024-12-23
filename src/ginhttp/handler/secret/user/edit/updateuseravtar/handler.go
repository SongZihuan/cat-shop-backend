package updateuseravtar

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
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
	CodeFileTooBig data.CodeType = -1
	CodeNotImage   data.CodeType = -2
)

func Handler(c *gin.Context) {
	user, ok := c.Value("User").(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	query := Query{}
	err := c.ShouldBindWith(&Query{}, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, data.NewClientBadRequests(err))
		return
	}

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

	err = action.UpdateUserAvatar(user, img.GetUrl())
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errImg))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
