package image

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
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
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	tp, ok := modeltype.NameToImageType[query.Type]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	hash := query.Hash
	if len(hash) != 64 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	basePath, ok := cfg.File.Image[tp]
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
	if !strings.HasPrefix("image/", mimeTp) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	acceptHeader := c.GetHeader(header.RequestsAccept)
	if acceptHeader != "" &&
		!strings.Contains(acceptHeader, "*/*") &&
		!strings.Contains(acceptHeader, "image/*") &&
		!strings.Contains(acceptHeader, mimeTp) {
		c.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set(header.RequestsContentType, mimeTp)
	_, _ = c.Writer.Write(dat)
}
