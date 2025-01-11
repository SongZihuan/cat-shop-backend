package config

import (
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"github.com/gin-gonic/gin"
	"os"
)

type LoggerLevel string

var levelMap = map[string]bool{
	"debug": true,
	"info":  true,
	"warn":  true,
	"error": true,
	"panic": true,
	"none":  true,
}

type GlobalConfig struct {
	Mode     string           `json:"mode"`
	LogLevel string           `json:"loglevel"`
	LogTag   utils.StringBool `json:"logtag"`
}

func (g *GlobalConfig) setDefault() {
	if g.Mode == "" {
		g.Mode = os.Getenv(gin.EnvGinMode)
	}

	if g.Mode == "" {
		g.Mode = gin.DebugMode
	}

	_ = os.Setenv(gin.EnvGinMode, g.Mode)

	if g.LogLevel == "" && (g.Mode == gin.DebugMode || g.Mode == gin.TestMode) {
		g.LogLevel = "debug"
	} else if g.LogLevel == "" {
		g.LogLevel = "warn"
	}

	if g.Mode == gin.DebugMode || g.Mode == gin.TestMode {
		g.LogTag.SetDefaultEnable()
	} else if g.LogLevel == "" {
		g.LogTag.SetDefaultDisable()
	}

	return
}

func (g *GlobalConfig) check() ConfigError {
	if g.Mode != gin.DebugMode && g.Mode != gin.ReleaseMode && g.Mode != gin.TestMode {
		return NewConfigError("bad mode")
	}

	if _, ok := levelMap[g.LogLevel]; !ok {
		return NewConfigError("log level error")
	}

	return nil
}

func (g *GlobalConfig) GetGinMode() string {
	return g.Mode
}

func (g *GlobalConfig) IsDebug() bool {
	return g.Mode == gin.DebugMode
}

func (g *GlobalConfig) IsRelease() bool {
	return g.Mode == gin.ReleaseMode
}

func (g *GlobalConfig) IsTest() bool {
	return g.Mode == gin.TestMode
}
