package loadpath

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/handler/v1/ping"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"net/http"
)

var pingHelloPath = ""

func LoadPingHelloPath(engine *ginplus.Router) {
	cfg := config.Config().Yaml.Http

	path, ok := engine.FindURLByHandler(ping.Handler, http.MethodGet)
	if !ok {
		pingPath = GetPingPath()
		if pingPath == "" {
			path = cfg.BasePath + cfg.PingPath + "/v1/helo"
		} else {
			path = pingPath + "/v1/hello"
		}
	}

	pingHelloPath = utils.ProcessPath(path)
	engine.DebugMsg("[INFO] ping hello path: %s", path)
}

func GetPingHelloPath() string {
	return pingHelloPath
}
