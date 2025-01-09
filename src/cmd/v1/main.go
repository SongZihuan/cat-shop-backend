package main

import (
	"github.com/SongZihuan/cat-shop-backend/src/cmd/define"
	"github.com/SongZihuan/cat-shop-backend/src/mainfunc"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

var v1Main define.MainFunc = mainfunc.MainV1

func main() {
	utils.Exit(_main())
}

func _main() int {
	utils.SayHello("shop backend server start")
	defer func() {
		utils.SayGoodBy("shop backend server stop")
	}()
	return v1Main()
}
