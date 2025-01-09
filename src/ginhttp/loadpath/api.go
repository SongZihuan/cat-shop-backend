package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/groupname"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

var apiPath = ""

func LoadAPIPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindGroupURLByName(groupname.APIName)
	if !ok {
		path = cfg.BasePath + cfg.ApiPath
	}

	apiPath = utils.ProcessPath(path)
	engine.DebugMsg("[INFO] api path: %s", path)
}

func GetAPIPath() string {
	return apiPath
}
