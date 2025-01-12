package adminrestartserver

import "github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"

type Query struct {
	Password string `json:"password"`
	Secret   string `json:"secret"`
	WaitSec  int    `json:"waitsec"`
}

type Data struct {
	Success bool `json:"success"`
	WaitSec int  `json:"waitsec"`
}

func NewSuccessData(waitsec int) Data {
	return Data{
		Success: true,
		WaitSec: waitsec,
	}
}

func NewSuccessJsonData(waitsec int) data.Data {
	return data.NewSuccessWithData(NewSuccessData(waitsec))
}
