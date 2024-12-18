package main

import (
	"github.com/SuperH-0630/cat-shop-back/src/cmd"
	"github.com/SuperH-0630/cat-shop-back/src/mainfunc"
)

var v1Main cmd.MainFunc = mainfunc.MainV1

func main() {
	cmd.Exit(_main())
}

func _main() int {
	cmd.SayHello("shop backend server start")
	defer func() {
		cmd.SayGoodBy("shop backend server stop")
	}()
	return v1Main()
}
