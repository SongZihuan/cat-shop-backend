package config

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type GlobalConfig struct {
	Mode string `json:"mode"`
}

func (g *GlobalConfig) setDefault() {
	if g.Mode == "" {
		g.Mode = "debug"
	}

	strings.ToLower(g.Mode)
	return
}

func (g *GlobalConfig) check() ConfigError {
	if g.Mode != "debug" && g.Mode != "release" && g.Mode != "test" {
		return NewConfigError("bad mode")
	}

	return nil
}

func (g *GlobalConfig) GetGinMode() string {
	switch g.Mode {
	case "debug":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
