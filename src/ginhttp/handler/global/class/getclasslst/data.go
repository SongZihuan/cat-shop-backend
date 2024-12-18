package getclasslst

import "github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"

type Data struct {
}

func NewData() Data {
	return Data{}
}

func NewJsonData() data.Data {
	return data.NewData(data.GlobalCodeOk, NewData())
}
