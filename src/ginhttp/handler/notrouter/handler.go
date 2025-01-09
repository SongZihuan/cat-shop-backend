package notrouter

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/get"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/options"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/other"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter/post"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler404(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodOptions:
		options.Handler(c)
	case http.MethodGet:
		get.Handler404(c)
	case http.MethodPost:
		post.Handler404(c)
	default:
		other.Handler404(c)
	}
}

func Handler403(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodOptions:
		options.Handler(c)
	case http.MethodGet:
		get.Handler403(c)
	case http.MethodPost:
		post.Handler403(c)
	default:
		other.Handler403(c)
	}
}
