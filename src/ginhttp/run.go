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

	_engine := gin.Default()

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
