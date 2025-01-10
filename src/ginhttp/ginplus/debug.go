package ginplus

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
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

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	_, _ = fmt.Fprintf(gin.DefaultWriter, "[GIN-PLUS-debug] "+format, values...)
}
