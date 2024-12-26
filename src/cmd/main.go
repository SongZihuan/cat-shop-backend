package cmd

import (
	"fmt"
	"os"
)

var args0 = ""

func init() {
	var err error
	if len(os.Args) > 0 {
		args0, err = os.Executable()
		if err != nil {
			args0 = os.Args[0]
		}
	}

	if args0 == "" {
		panic("args was empty")
	}
}

type MainFunc func() int

func Exit(code ...int) {
	if len(code) == 1 {
		os.Exit(code[0])
	}
	os.Exit(0)
}

func SayHello(msg ...string) {
	if len(msg) == 1 {
		fmt.Println(args0, ":", msg[0])
	}
	fmt.Println(args0, ":", "start")
}

func SayGoodBy(msg ...string) {
	if len(msg) == 1 {
		fmt.Println(args0, ":", msg[0])
	}
	fmt.Println(args0, ":", "end")
}
