package ginhttp

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/router"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine = nil

func InitGin() error {
	if !config.IsReady() {
		panic("config is not ready")
	}
	cfg := config.Config()

	gin.SetMode(cfg.Yaml.Global.GetGinMode())

	_engine := gin.Default()

	if cfg.Yaml.Http.Proxy.Enable() {
		_engine.ForwardedByClientIP = true
		err := _engine.SetTrustedProxies(cfg.Yaml.Http.Proxy.TrustedIPs)
		if err != nil {
			return err
		}
	}

	router.InitRouter(_engine)

	engine = _engine
	return nil
}

func Run() error {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if engine == nil {
		return fmt.Errorf("router is nil")
	}

	err := engine.Run(config.Config().Yaml.Http.Address)
	if err != nil {
		return err
	}

	return nil
}
