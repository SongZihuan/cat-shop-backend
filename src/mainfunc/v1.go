package mainfunc

import (
	"errors"
	"flag"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/flagparser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/httpstop"
	"time"
)

func MainV1() int {
	var err error

	err = flagparser.Flag()
	if errors.Is(err, flag.ErrHelp) {
		return 0
	} else if err != nil {
		exitByError(err)
		return 1
	}

	if !flagparser.IsReady() {
		exitByMsg("flag parser unknown error")
		return 1
	}

	err = config.InitConfig()
	if err != nil {
		exitByError(err)
		return 1
	}

	if !config.IsReady() {
		exitByMsg("config parser unknown error")
		return 1
	}

	cfg := config.Config()

	if flagparser.Wait() {
		<-time.Tick(time.Duration(cfg.Yaml.Http.RestartWaitSecond) * time.Second)
	}

	err = database.ConnectToMySql()
	if err != nil {
		exitByError(err)
		return 1
	}
	defer database.CloseMySql()

	err = action.AutoMigrate()
	if err != nil {
		exitByError(err)
		return 1
	}

	err = action.CreateEmptyClass()
	if err != nil {
		exitByError(err)
		return 1
	}

	err = ginhttp.InitGin()
	if err != nil {
		exitByError(err)
		return 1
	}

	fmt.Printf("run mode: %s\n", cfg.Yaml.Global.GetGinMode())

	ginstop := make(chan bool)
	ginerror := make(chan error)

	go func() {
		err = ginhttp.Run()
		if errors.Is(err, ginhttp.ServerClose) {
			ginstop <- true
		} else if err != nil {
			ginerror <- err
		} else {
			ginstop <- true
		}
	}()

	select {
	case <-cfg.GetSignalChan():
		break
	case <-httpstop.GetStopChan():
		break
	case err := <-ginerror:
		exitByError(err)
		return 1
	case <-ginstop:
		break
	}

	err = ginhttp.Stop(time.Duration(cfg.Yaml.Http.StopWaitSecond) * time.Second)
	if err != nil {
		exitByError(err)
		return 1
	}

	return 0
}
