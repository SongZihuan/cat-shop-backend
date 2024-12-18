package cmd

import (
	"fmt"
	"os"
)

var args0 = ""

func init() {
	if len(os.Args) > 0 {
		args0 = os.Args[0]
	} else {
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
		fmt.Println(args0, ": ", msg)
	}
	fmt.Println(args0, ": start.")
}

func SayGoodBy(msg ...string) {
	if len(msg) == 1 {
		fmt.Println(args0, ": ", msg)
	}
	fmt.Println(args0, ": end.")
}
