package video

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/abort"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

func Handler(c *gin.Context) {
	cfg := config.Config()

	query := Query{}
	err := c.ShouldBindQuery(&Query{})
	if err != nil {
		abort.BadRequestsError(c, err)
		return
	}

	tp, ok := modeltype.NameToVideoType[query.Type]
	if !ok {
		abort.ResourceNotFound(c)
		return
	}

	hash := query.Hash
	if len(hash) != 64 {
		abort.ResourceNotFound(c)
		return
	}

	basePath, ok := cfg.File.Video[tp]
	if !ok {
		abort.ResourceNotFound(c)
		return
	}

	ph := path.Join(basePath, query.Time, fmt.Sprintf("%s.dat", hash))
	dat, err := os.ReadFile(ph)
	if err != nil {
		abort.ResourceNotFound(c)
		return
	}

	mimeTp := mimetype.Detect(dat).String()
	if !strings.HasPrefix("video/", mimeTp) {
		abort.ResourceNotFound(c)
		return
	}

	acceptHeader := c.GetHeader(header.RequestsAccept)
	if acceptHeader != "" &&
		!strings.Contains(acceptHeader, "*/*") &&
		!strings.Contains(acceptHeader, "video/*") &&
		!strings.Contains(acceptHeader, mimeTp) {
		abort.ResourceNotAccept(c, fmt.Sprintf("*/*,image/*,%s", mimeTp))
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set(header.RequestsContentType, mimeTp)
	_, _ = c.Writer.Write(dat)
}
