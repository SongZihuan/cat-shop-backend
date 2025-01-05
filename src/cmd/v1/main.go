package main

import (
	"github.com/SongZihuan/cat-shop-backend/src/cmd"
	"github.com/SongZihuan/cat-shop-backend/src/mainfunc"
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
