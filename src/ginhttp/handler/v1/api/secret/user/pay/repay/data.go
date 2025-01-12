package repay

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
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
	return data.NewSuccessWithData(NewData(url))
}
