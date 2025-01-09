package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/groupname"
)

var apiPath = ""

func LoadAPIPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindGroupURLByName(groupname.APIName)
	if !ok {
		path = cfg.BasePath + cfg.ApiPath
	}

	videosPath = path
	engine.DebugMsg("[INFO] api path: %s", path)
}

func GetAPIPath() string {
	return apiPath
}
