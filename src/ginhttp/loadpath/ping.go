package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router/groupname"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

var pingPath = ""

func LoadPingPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindGroupURLByName(groupname.PingName)
	if !ok {
		path = cfg.BasePath + cfg.PingPath
	}

	pingPath = utils.ProcessPath(path)
	engine.DebugMsg("[INFO] ping path: %s", path)
}

func GetPingPath() string {
	return pingPath
}
