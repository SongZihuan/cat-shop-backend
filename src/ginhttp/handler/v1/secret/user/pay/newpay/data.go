package newpay

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

type Query struct {
	Type       modeltype.PayFromType `form:"type"`
	RedirectTo string                `form:"redirectto"`
	WupinID    uint                  `form:"wupinId"`
	Num        modeltype.Total       `form:"num"`

	UserName     string `form:"username"`
	UserPhone    string `form:"userphone"`
	UserLocation string `form:"userlocation"`
	UserWechat   string `form:"userwechat"`
	UserEmail    string `form:"useremail"`
	UserRemark   string `form:"userremark"`
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
