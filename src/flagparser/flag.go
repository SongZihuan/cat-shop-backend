package flagparser

import (
	"flag"
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

var isReady = false

func IsReady() bool {
	return data.isReady() && isReady
}

var StopFlag = fmt.Errorf("stop")

func InitFlag() (err error) {
	if isReady {
		return nil
	}

	defer func() {
		if e := recover(); e != nil {
			err = NewFlagError(e)
			return
		}
	}()

	initData()

	if Version() {
		_, _ = FprintVersion(flag.CommandLine.Output())
	}

	if License() {
		_, _ = FprintLicense(flag.CommandLine.Output())
	}

	if Help() {
		_, _ = FprintUseage(flag.CommandLine.Output())
	}

	if NotRunMode() {
		return StopFlag
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
