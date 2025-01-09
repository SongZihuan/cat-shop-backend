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

func Exit(code ...int) {
	if len(code) == 1 {
		os.Exit(code[0])
	}
	os.Exit(0)
}

func SayHello(msg ...string) {
	if len(msg) == 1 {
		fmt.Printf("%s: %s\n", GetArgs0Name(), msg[0])
	}
	fmt.Printf("%s: %s\n", GetArgs0Name(), "start to run")
}

func SayGoodBy(msg ...string) {
	if len(msg) == 1 {
		fmt.Printf("%s: %s\n", GetArgs0Name(), msg[0])
	}
	fmt.Printf("%s: %s\n", GetArgs0Name(), "stop running")
}
