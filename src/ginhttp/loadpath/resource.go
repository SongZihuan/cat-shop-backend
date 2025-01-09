package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/groupname"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

var resourcePath = ""

func LoadResourcePath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindGroupURLByName(groupname.ResourceName)
	if !ok {
		path = cfg.BasePath + cfg.ResourcePath
	}

	resourcePath = utils.ProcessPath(path)
	engine.DebugMsg("[INFO] resource path: %s", path)
}

func GetResourcePath() string {
	return resourcePath
}
