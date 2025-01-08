package router

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/middleware"
	v1 "github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/v1"
)

const Version1 = "/v1"

func InitRouter(engine *ginplus.Router) {
	v1Router(engine)
}

func v1Router(engine *ginplus.Router) {
	if !config.IsReady() {
		panic("config is not ready")
	}

	engine.Use(middleware.AllReady())
	v1.Api(engine.Group(config.Config().Yaml.Http.BaseURL).Group(config.Config().Yaml.Http.ApiURL).Group(Version1))
	v1.Resource(engine.Group(config.Config().Yaml.Http.BaseURL).Group(config.Config().Yaml.Http.ResourceURL).Group(Version1))
}
