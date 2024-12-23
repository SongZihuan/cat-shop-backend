package repay

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

type Query struct {
	ID         uint                  `form:"id"`
	Type       modeltype.PayFromType `form:"type"`
	RedirectTo string                `form:"redirectto"`
}

type Data struct {
	Url string `json:"url"`
}

func NewData(url string) Data {
	return Data{
		Url: url,
	}
}

func NewJsonData(url string) data.Data {
	return data.NewData(data.GlobalCodeOk, NewData(url))
}
