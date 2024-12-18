package mainfunc

import (
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/database"
	"github.com/SuperH-0630/cat-shop-back/src/database/action"
	"github.com/SuperH-0630/cat-shop-back/src/flagparser"
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp"
)

func MainV1() int {
	var err error

	err = flagparser.Flag()
	if err != nil {
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

	err = ginhttp.InitGin()
	if err != nil {
		exitByError(err)
		return 1
	}

	err = ginhttp.Run()
	if err != nil {
		exitByError(err)
		return 1
	}

	return 0
}
