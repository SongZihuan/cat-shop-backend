package video

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model"
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
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	tp, ok := model.NameToVideoType[query.Type]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	hash := query.Hash
	if len(hash) != 64 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	basePath, ok := cfg.File.Video[tp]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	ph := path.Join(basePath, query.Time, fmt.Sprintf("%s.dat", hash))
	dat, err := os.ReadFile(ph)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	mimeTp := mimetype.Detect(dat).String()
	if !strings.HasPrefix("video/", mimeTp) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	acceptHeader := c.GetHeader("Accept")
	if acceptHeader != "" &&
		!strings.Contains(acceptHeader, "*/*") &&
		!strings.Contains(acceptHeader, "video/*") &&
		!strings.Contains(acceptHeader, mimeTp) {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", mimeTp)
	_, _ = c.Writer.Write(dat)
}
