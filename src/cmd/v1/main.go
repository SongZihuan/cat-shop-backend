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
	utils.SayHellof("%s", "The backend service starts normally, thank you.")
	defer func() {
		utils.SayGoodByef("%s", "The backend service is offline/shutdown normally, thank you.")
	}()
	return v1Main()
}
