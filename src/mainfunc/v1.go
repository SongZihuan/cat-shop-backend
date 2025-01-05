package mainfunc

import (
	"errors"
	"flag"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/database"
	"github.com/SongZihuan/cat-shop-backend/src/database/action"
	"github.com/SongZihuan/cat-shop-backend/src/flagparser"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/httpstop"
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

	waitsec := flagparser.WaitSec()
	if waitsec > 0 {
		time.Sleep(waitsec)
	}

	err = database.ConnectToMysql()
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
