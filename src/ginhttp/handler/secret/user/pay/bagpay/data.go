package bagpay

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
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
}

func NewData() Data {
	return Data{}
}

func NewJsonData() data.Data {
	return data.NewData(data.GlobalCodeOk, NewData())
}
