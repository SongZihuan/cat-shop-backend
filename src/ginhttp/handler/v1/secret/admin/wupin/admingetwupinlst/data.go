package admingetwupinlst

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	Page     int `form:"page"`
	PageSize int `form:"pagesize"`
}

type Data struct {
	List     []Wupin `json:"list"`
	Total    int     `json:"total"`
	MaxCount int     `json:"maxpage"`
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

func NewWupin(wupin model.Wupin) Wupin {
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

func NewData(list []model.Wupin, maxcount int) Data {
	res := make([]Wupin, 0, len(list))
	for _, v := range list {
		res = append(res, NewWupin(v))
	}

	return Data{
		List:     res,
		Total:    len(list),
		MaxCount: maxcount,
	}
}

func NewJsonData(list []model.Wupin, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(list, maxcount))
}
