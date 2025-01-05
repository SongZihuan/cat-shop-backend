package flagparser

import (
	"flag"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
)

var isReady = false

func IsReady() bool {
	return data.isReady() && isReady
}

var ErrHelp = fmt.Errorf("help")

func Flag() (err error) {
	if isReady {
		return nil
	}

	defer func() {
		if e := recover(); e != nil {
			err = NewFlagError(e)
			return
		}
	}()

	data = newFlagData()

	flag.BoolVar(&data.help, "help", false, "this help")
	flag.StringVar(&data.configFile, "config", "config.yaml", "the config file path")
	flag.UintVar(&data.wait, "wait", MinWaitSec, "wait second to start")

	flag.Parse()
	data.ready()

	if Help() {
		flag.Usage()
		return ErrHelp
	}

	err = checkFlag()
	if err != nil {
		return err
	}

	isReady = true
	return nil
}

func checkFlag() error {
	if !utils.IsExists(ConfigFile()) {
		return fmt.Errorf("config file not exists")
	}

	return nil
}
