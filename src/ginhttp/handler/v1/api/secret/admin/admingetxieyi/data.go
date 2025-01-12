package admingetxieyi

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type Query struct {
	Type modeltype.XieYiType `json:"type"`
}

type Data struct {
	Xieyi string `json:"xieyi"`
}

func NewData(xieyi string) Data {
	return Data{
		Xieyi: xieyi,
	}
}

func NewJsonData(xieyi string) data.Data {
	return data.NewSuccessWithData(NewData(xieyi))
}
