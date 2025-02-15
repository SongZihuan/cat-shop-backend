package bagpay

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type Query struct {
	Type       modeltype.PayFromType `form:"type"`
	RedirectTo string                `form:"redirectto"`
	BagID      uint                  `form:"bagid"`

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
