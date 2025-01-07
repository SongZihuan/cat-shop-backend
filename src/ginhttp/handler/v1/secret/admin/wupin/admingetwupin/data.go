package admingetwupin

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	ID uint `json:"id"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Show bool   `json:"show"`
	Down bool   `json:"down"`
}

type Wupin struct {
	ID         uint                `json:"id"`
	Name       string              `json:"name"`
	ClassID    uint                `json:"classId"`
	Tag        string              `json:"tag"`
	HotPrice   modeltype.PriceNull `json:"hotPrice"`
	RealPrice  modeltype.Price     `json:"realPrice"`
	Info       string              `json:"info"`
	Ren        string              `json:"ren"`
	Phone      string              `json:"phone"`
	Email      string              `json:"email"`
	Wechat     string              `json:"wechat"`
	Location   string              `json:"location"`
	Pic        string              `json:"pic"`
	ClassOf    *Class              `json:"classOf"`
	BuyTotal   modeltype.Total     `json:"buytotal"`
	BuyDaohuo  modeltype.Total     `json:"buydaohuo"`
	BuyGood    modeltype.Total     `json:"buygood"`
	BuyPrice   modeltype.Price     `json:"buyprice"`
	BuyPingjia modeltype.Total     `json:"buypingjia"`
	BuyJian    modeltype.Total     `json:"buyjian"`
	Hot        bool                `json:"hot"`
	Down       bool                `json:"down"`
}

func NewData(wupin *model.Wupin) Wupin {
	return Wupin{
		ID:        wupin.ID,
		Name:      wupin.Name,
		ClassID:   wupin.ClassID,
		Tag:       utils.GetSQLNullString(wupin.Tag),
		HotPrice:  wupin.HotPrice,
		RealPrice: wupin.RealPrice,
		Info:      wupin.Info,
		Ren:       wupin.Ren,
		Phone:     wupin.Phone,
		Email:     utils.GetSQLNullString(wupin.Email),
		Wechat:    utils.GetSQLNullString(wupin.WeChat),
		Location:  wupin.Location,
		Pic:       wupin.Pic,
		ClassOf: &Class{
			ID:   wupin.ClassID,
			Name: wupin.Class.Name,
			Show: wupin.Class.IsClassShow(),
			Down: wupin.Class.IsClassDown(),
		}, // Assuming ClassID matches Class.ID
		BuyTotal:   wupin.BuyTotal,   // 购物总人数
		BuyDaohuo:  wupin.BuyDaoHuo,  // 到货总人数
		BuyGood:    wupin.BuyGood,    // 好评总人数
		BuyPrice:   wupin.BuyMoney,   // 购物总金额
		BuyPingjia: wupin.BuyPingjia, // 评价总人数
		BuyJian:    wupin.BuyJian,    // 购买总件数
		Hot:        wupin.Hot,
		Down:       wupin.IsWupinDown(),
	}
}

func NewJsonData(wupin *model.Wupin) data.Data {
	return data.NewSuccessWithData(NewData(wupin))
}
