package ginplus

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/logger"
	"github.com/gin-gonic/gin"
	"strings"
)

func isDebugging() bool {
	if !config.IsReady() {
		panic("config is not ready")
	}

	return config.Config().Yaml.GlobalConfig.IsDebug()
}

func debugPrint(format string, values ...any) {
	if !isDebugging() {
		return
	}

	if gin.DebugPrintFunc != nil {
		gin.DebugPrintFunc(format, values...)
		return
	}

	res := fmt.Sprintf(format, values...)
	res = strings.TrimRight(res, "\n")
	logger.Debugf("[GIN-PLUS-debug] %s", res)
}
