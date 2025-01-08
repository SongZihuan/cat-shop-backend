package cors

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/header"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Handler(c *gin.Context) bool {
	origin := c.GetHeader(header.RequestsOrigin)
	if origin == "" {
		c.JSON(http.StatusMethodNotAllowed, data.NewClientCorsError("没用Origin头部"))
		return false
	}

	if !config.Config().CoreOrigin.InOriginList(origin) {
		c.JSON(http.StatusMethodNotAllowed, data.NewClientCorsError("Origin头部不符合规定"))
		return false
	}

	c.Header(header.ResponseAllowOrigin, origin)
	c.Header(header.ResponseAllowMethods, strings.Join(header.AllowMethods, ","))
	c.Header(header.ResponseExposeHeaders, strings.Join(header.RequestsHeaderList, ","))
	c.Header(header.ResponseExposeHeaders, strings.Join(header.ResponseHeaderList, ","))
	c.Header(header.ResponseAllowCredentials, "false")
	c.Header(header.ResponseAllowMaxAge, fmt.Sprintf("%d", config.Config().Yaml.Http.Cors.MaxAgeSec))

	return true
}
