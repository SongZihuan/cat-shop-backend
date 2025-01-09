package mainfunc

import (
	"errors"
	"flag"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database"
	"github.com/SongZihuan/cat-shop-backend/src/database/action/automigrator"
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/httpstop"
	"github.com/SongZihuan/cat-shop-backend/src/logger"
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

	err = logger.InitLogger()
	if err != nil {
		exitByError(err)
		return 1
	}

	if !logger.IsReady() {
		exitByMsg("logger unknown error")
		return 1
	}

	waitsec := flagparser.WaitSec()
	if waitsec > 0 {
		time.Sleep(waitsec)
	}

	err = database.ConnectToMySQL()
	if err != nil {
		exitByError(err)
		return 1
	}
	defer database.CloseMySQL()

	err = automigrator.SystemAutoMigrate()
	if err != nil {
		exitByError(err)
		return 1
	}

	err = automigrator.SystemCreateEmptyClass()
	if err != nil {
		exitByError(err)
		return 1
	}

	err = ginhttp.InitGin()
	if err != nil {
		exitByError(err)
		return 1
	}

	logger.Infof("run mode: %s\n", cfg.Yaml.Global.GetGinMode())

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
