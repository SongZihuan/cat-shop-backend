package ginhttp

import (
	"context"
	"errors"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/ginplus"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/resourcepath"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/router"
	"net/http"
	"time"
)

var ServerClose = fmt.Errorf("server close")

var engine *ginplus.Router = nil
var server *http.Server = nil

func InitGin() error {
	_engine, err := ginplus.NewEngine()
	if err != nil {
		return err
	}
	router.InitRouter(_engine)
	resourcepath.LoadImagePath(_engine)
	resourcepath.LoadVideoPath(_engine)
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
