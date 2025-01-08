package ginplus

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
)

func IsDebugging() bool {
	if !config.IsReady() {
		panic("config is not ready")
	}

	return config.Config().Yaml.Global.IsDebug()
}
