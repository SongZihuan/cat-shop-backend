package getxieyi

import "github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"

type Data struct {
	XieYi string `json:"xieyi"`
}

func NewData(XieYi string) Data {
	return Data{XieYi: XieYi}
}

func NewJsonData(XieYi string) data.Data {
	return data.NewSuccessWithData(NewData(XieYi))
}
