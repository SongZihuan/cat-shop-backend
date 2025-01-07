package getbaglst

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type WuPin struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Pic        string `json:"pic"`
	ClassID    uint   `json:"classId"`
	ClassOf    *Class `json:"classOf"`
	Tag        string `json:"tag,omitempty"`
	HotPrice   int64  `json:"hotPrice,omitempty"`
	RealPrice  int64  `json:"realPrice"`
	Info       string `json:"info"`
	Ren        string `json:"ren"`
	Phone      string `json:"phone"`
	Email      string `json:"email,omitempty"`
	Wechat     string `json:"wechat,omitempty"`
	Location   string `json:"location"`
	BuyTotal   int64  `json:"buytotal"`
	BuyDaohuo  int64  `json:"buydaohuo"`
	BuyGood    int64  `json:"buygood"`
	BuyPrice   int64  `json:"buyprice"`
	BuyPingjia int64  `json:"buypingjia"`
	BuyJian    int64  `json:"buyjian"`
	Hot        bool   `json:"hot"`
}

type Bag struct {
	ID      uint            `json:"id"`
	UserID  uint            `json:"userId"`
	WupinID uint            `json:"wupinId"`
	ClassID uint            `json:"classId"`
	Num     modeltype.Total `json:"num"`
	Time    int64           `json:"time"`
	Wupin   *WuPin          `json:"wupin"`
	Down    bool            `json:"down"`
}

func NewBag(bag *model.Bag) Bag {
	if bag.IsBagCanSale() {
		return Bag{
			ID:      bag.ID,
			UserID:  bag.UserID,
			WupinID: bag.WupinID,
			ClassID: bag.ClassID,
			Num:     bag.Num,
			Time:    bag.Time.Unix(),
			Wupin: &WuPin{
				ID:      bag.Wupin.ID,
				Name:    bag.Wupin.Name,
				Pic:     bag.Wupin.Pic,
				ClassID: bag.ClassID,
				ClassOf: &Class{
					ID:   bag.Class.ID,
					Name: bag.Class.Name,
				},
				Tag:        utils.GetSQLNullString(bag.Wupin.Tag),
				HotPrice:   bag.Wupin.HotPrice.ToInt64(),
				RealPrice:  bag.Wupin.RealPrice.ToInt64(),
				Info:       bag.Wupin.Info,
				Ren:        bag.Wupin.Ren,
				Phone:      bag.Wupin.Phone,
				Email:      utils.GetSQLNullString(bag.Wupin.Email),
				Wechat:     utils.GetSQLNullString(bag.Wupin.WeChat),
				Location:   bag.Wupin.Location,
				BuyTotal:   bag.Wupin.BuyTotal.ToInt64(),
				BuyDaohuo:  bag.Wupin.BuyDaoHuo.ToInt64(),
				BuyGood:    bag.Wupin.BuyGood.ToInt64(),
				BuyPrice:   bag.Wupin.BuyPrice.ToInt64(),
				BuyPingjia: bag.Wupin.BuyPingjia.ToInt64(),
				BuyJian:    bag.Wupin.BuyJian.ToInt64(),
				Hot:        bag.Wupin.Hot,
			},
			Down: false,
		}
	} else {
		return Bag{
			ID:      bag.ID,
			UserID:  bag.UserID,
			WupinID: bag.WupinID,
			ClassID: bag.ClassID,
			Num:     bag.Num,
			Time:    bag.Time.Unix(),
			Wupin: &WuPin{
				ID:      bag.Wupin.ID,
				Name:    fmt.Sprintf("%s（以下架）", bag.Wupin.Name),
				Pic:     bag.Wupin.Pic,
				ClassID: bag.ClassID,
				ClassOf: &Class{
					ID:   bag.Class.ID,
					Name: bag.Class.Name,
				},
				Tag:        "",
				HotPrice:   0,
				RealPrice:  0,
				Info:       bag.Wupin.Info,
				Ren:        "无",
				Phone:      "000-0000-0000",
				Email:      "",
				Wechat:     "",
				Location:   "无",
				BuyTotal:   bag.Wupin.BuyTotal.ToInt64(),
				BuyDaohuo:  bag.Wupin.BuyDaoHuo.ToInt64(),
				BuyGood:    bag.Wupin.BuyGood.ToInt64(),
				BuyPrice:   bag.Wupin.BuyPrice.ToInt64(),
				BuyPingjia: bag.Wupin.BuyPingjia.ToInt64(),
				BuyJian:    bag.Wupin.BuyJian.ToInt64(),
				Hot:        false,
			},
			Down: true,
		}
	}
}

type Data struct {
	List  []Bag `json:"list"`
	Total int   `json:"total"`
}

func NewData(res []model.Bag) Data {
	list := make([]Bag, len(res))
	for _, v := range res {
		if v.IsBagShow() {
			list = append(list, NewBag(&v))
		}
	}

	return Data{
		List:  list,
		Total: len(res),
	}
}

func NewJsonData(res []model.Bag) data.Data {
	return data.NewSuccessWithData(NewData(res))
}
