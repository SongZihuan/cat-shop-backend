package ginhttp

import (
	"context"
	"errors"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var ServerClose = fmt.Errorf("server close")

var engine *gin.Engine = nil
var server *http.Server = nil

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

	server = &http.Server{
		Addr:    config.Config().Yaml.Http.Address,
		Handler: engine,
	}

	err := server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return ServerClose
		}
		return err
	}

	return nil
}

func Stop(waitTime ...time.Duration) error {
	var ctx context.Context
	var cancel context.CancelFunc

	if len(waitTime) == 0 {
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	} else if len(waitTime) == 1 {
		if waitTime[0] == 0 {
			waitTime[0] = 5 * time.Second
		}
		ctx, cancel = context.WithTimeout(context.Background(), waitTime[0])
		defer cancel()
	} else {
		panic("too many arguments")
	}

	return server.Shutdown(ctx)
}
