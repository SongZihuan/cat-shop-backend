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
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"time"
)

func MainV1() int {
	var err error

	err = flagparser.Flag()
	if errors.Is(err, flag.ErrHelp) {
		return 0
	} else if err != nil {
		return utils.ExitByError(err)
	}

	if !flagparser.IsReady() {
		return utils.ExitByErrorMsg("flag parser unknown error")
	}

	err = config.InitConfig()
	if err != nil {
		return utils.ExitByError(err)
	}

	if !config.IsReady() {
		return utils.ExitByErrorMsg("config parser unknown error")
	}

	cfg := config.Config()

	err = logger.InitLogger()
	if err != nil {
		return utils.ExitByError(err)
	}

	if !logger.IsReady() {
		return utils.ExitByErrorMsg("logger unknown error")
	}

	waitsec := flagparser.WaitSec()
	if waitsec > 0 {
		time.Sleep(waitsec)
	}

	err = database.ConnectToMySQL()
	if err != nil {
		return utils.ExitByError(err)
	}
	defer database.CloseMySQL()

	err = automigrator.SystemAutoMigrate()
	if err != nil {
		return utils.ExitByError(err)
	}

	err = automigrator.SystemCreateEmptyClass()
	if err != nil {
		return utils.ExitByError(err)
	}

	err = ginhttp.InitGin()
	if err != nil {
		return utils.ExitByError(err)
	}

	logger.Executable()
	logger.Infof("run mode: %s\n", cfg.Yaml.GlobalConfig.GetGinMode())

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
		return utils.ExitByError(err)
	case <-ginstop:
		break
	}

	err = ginhttp.Stop(time.Duration(cfg.Yaml.Http.StopWaitSecond) * time.Second)
	if err != nil {
		return utils.ExitByError(err)
	}

	return 0
}
