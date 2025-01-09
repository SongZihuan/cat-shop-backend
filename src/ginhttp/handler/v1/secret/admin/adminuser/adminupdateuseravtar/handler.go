package adminupdateuseravtar

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/adminaction"
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
	CodeFileTooBig    data.CodeType = -4
	CodeNotImage      data.CodeType = -5
	CodeUserIsDelete  data.CodeType = -6
	CodeFileNotUpload data.CodeType = -6
)

func Handler(c *gin.Context) {
	user, ok := c.Value(contextkey.AdminUserKey).(*model.User)
	if !ok {
		c.JSON(http.StatusOK, data.NewSystemUnknownError("用户未找到"))
		return
	}

	if err := c.Request.ParseMultipartForm(Size3MB); err != nil { // 32MB限制
		abort.BadRequestsError(c, err)
		return
	}

	if user.IsDeleteUser() {
		c.JSON(http.StatusOK, data.NewCustomError(CodeUserIsDelete, "用户已经被删除")) // 已经删除是用户无法执行操作
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

	img, errDB, errImg := adminaction.AdminNewImage(modeltype.AvatarImage, fileData)
	if errImg != nil {
		c.JSON(http.StatusOK, data.NewSystemUnknownError(errImg))
		return
	} else if errDB != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errDB))
		return
	}

	err = adminaction.AdminUpdateUserAvatar(user, img.GetUrl())
	if err != nil {
		c.JSON(http.StatusOK, data.NewSystemDataBaseError(errImg))
		return
	}

	c.JSON(http.StatusOK, data.NewSuccess("更新成功"))
}
