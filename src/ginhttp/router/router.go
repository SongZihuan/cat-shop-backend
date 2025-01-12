package router

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/notrouter"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/middleware"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/groupname"
	v1 "github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/v1"
)

const Version1 = "/v1"

func InitRouter(engine *ginplus.Router) {
	engine.NotRouter(notrouter.Handler404, notrouter.Handler403)
	v1Router(engine)
}

func v1Router(engine *ginplus.Router) {
	if !config.IsReady() {
		panic("config is not ready")
	}

	engine.Use(middleware.AllReady(), middleware.AllowMethod(), middleware.Cors())
	v1.Ping(engine.Group(config.Config().Yaml.Http.BasePath).Group(config.Config().Yaml.Http.PingPath, groupname.PingName).Group(Version1))
	v1.Api(engine.Group(config.Config().Yaml.Http.BasePath).Group(config.Config().Yaml.Http.ApiPath, groupname.APIName).Group(Version1))
	v1.Resource(engine.Group(config.Config().Yaml.Http.BasePath).Group(config.Config().Yaml.Http.ResourcePath, groupname.ResourceName).Group(Version1))
}
