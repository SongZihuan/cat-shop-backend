package flagparser

import (
	"flag"
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"os"
)

var isReady = false

func IsReady() bool {
	return data.isReady() && isReady
}

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

	flag.Parse()
	data.ready()

	if Help() {
		flag.Usage()
		os.Exit(0)
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
