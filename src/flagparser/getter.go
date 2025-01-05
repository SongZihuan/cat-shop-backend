package flagparser

import "time"

var data flagData

func Help() bool {
	return data.Help()
}

func ConfigFile() string {
	return data.ConfigFile()
}

func WaitSec() time.Duration {
	return time.Second * time.Duration(data.Wait())
}
