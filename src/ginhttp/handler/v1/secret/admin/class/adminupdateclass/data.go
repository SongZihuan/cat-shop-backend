package adminupdateclass

import "github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"

type Data struct {
}

func NewData() Data {
	return Data{}
}

func NewJsonData() data.Data {
	return data.NewSuccessWithData(NewData())
}
