package getxieyi

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type Query struct {
	Type modeltype.XieYiType `json:"type"`
}

type Data struct {
	XieYi string `json:"xieyi"`
}

func NewData(XieYi string) Data {
	return Data{XieYi: XieYi}
}

func NewJsonData(XieYi string) data.Data {
	return data.NewSuccessWithData(NewData(XieYi))
}
