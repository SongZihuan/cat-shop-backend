package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var _args0 = ""

func init() {
	var err error
	if len(os.Args) > 0 {
		_args0, err = os.Executable()
		if err != nil {
			_args0 = os.Args[0]
		}
	}

	if _args0 == "" {
		panic("args was empty")
	}
}

func GetArgs0() string {
	return _args0
}

func GetArgs0Name() string {
	return filepath.Base(_args0)
}

func SayHellof(format string, args ...interface{}) {
	if len(format) == 0 && len(args) == 0 {
		fmt.Printf("%s: %s\n", GetArgs0Name(), "Normal startup, thank you.")
	} else {
		str := fmt.Sprintf(format, args...)
		fmt.Printf("%s: %s\n", GetArgs0Name(), str)
	}
}

func SayGoodByef(format string, args ...interface{}) {
	if len(format) == 0 && len(args) == 0 {
		fmt.Printf("%s: %s\n", GetArgs0Name(), "Normal shutdown, thank you.")
	} else {
		str := fmt.Sprintf(format, args...)
		fmt.Printf("%s: %s\n", GetArgs0Name(), str)
	}
}
