package mainfunc

import (
	"errors"
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

	err = flagparser.InitFlag()
	if errors.Is(err, flagparser.StopFlag) {
		return 0
	} else if err != nil {
		return utils.ExitByError(err)
	}

	if !flagparser.IsReady() {
		return utils.ExitByErrorMsg("flag parser unknown error")
	}

	utils.SayHellof("%s", "The backend service program starts normally, thank you.")
	defer func() {
		utils.SayGoodByef("%s", "The backend service program is offline/shutdown normally, thank you.")
	}()

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

	err = database.ConnectToMySQL()
	if err != nil {
		return utils.ExitByError(err)
	}
	defer database.CloseMySQL()

	err = automigrator.SystemAutoMigrate()
	if err != nil {
		return utils.ExitByError(err)
	}

	err = automigrator.SystemMustCreateData()
	if err != nil {
		return utils.ExitByError(err)
	}

	err = automigrator.SystemCreateData()
	if err != nil {
		return utils.ExitByError(err)
	}

	waitsec := flagparser.WaitSec()
	if waitsec > 0 {
		logger.Infof("The backend service process is sleeping and waiting for %d seconds.", waitsec)
		time.Sleep(waitsec)
		logger.Infof("%s", "Backend service sleeps and waits for completion")
	}

	logger.Executable()
	logger.Infof("run mode: %s\n", cfg.Yaml.GlobalConfig.GetGinMode())

	err = ginhttp.InitGin()
	if err != nil {
		return utils.ExitByError(err)
	}

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
