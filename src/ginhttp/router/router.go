package router

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	v1 "github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/v1"
	"github.com/gin-gonic/gin"
)

const VersionV1 = "/v1"

func InitRouter(engine *gin.Engine) {
	if !config.IsReady() {
		panic("config is not ready")
	}
	v1Router(engine)
}

func v1Router(engine *gin.Engine) {
	v1.Api(engine.Group(config.Config().Yaml.Http.ApiBaseAPI).Group(VersionV1))
	v1.Resource(engine.Group(config.Config().Yaml.Http.ResourceBaseAPI).Group(VersionV1))
}
