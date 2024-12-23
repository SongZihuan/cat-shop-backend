package getxieyi

import "github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"

type Data struct {
	XieYi string `json:"xieyi"`
}

func NewData(XieYi string) Data {
	return Data{XieYi: XieYi}
}

func NewJsonData(XieYi string) data.Data {
	return data.NewSuccessWithData(NewData(XieYi))
}
