package getbaglstbypage

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
)

type Query struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Wupin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	ClassID   uint   `json:"classId"`
	ClassOf   *Class `json:"classOf"`
	Tag       string `json:"tag,omitempty"`
	HotPrice  int64  `json:"hotPrice,omitempty"`
	RealPrice int64  `json:"realPrice"`
	Info      string `json:"info"`
	Ren       string `json:"ren"`
	Phone     string `json:"phone"`
	Email     string `json:"email,omitempty"`
	Wechat    string `json:"wechat,omitempty"`
	Location  string `json:"location"`
	BuyTotal  int64  `json:"buytotal"`
	BuyDaohuo int64  `json:"buydaohuo"`
	BuyGood   int64  `json:"buygood"`
}

type Bag struct {
	ID      uint            `json:"id"`
	UserID  uint            `json:"userId"`
	WupinID uint            `json:"wupinId"`
	ClassID uint            `json:"classId"`
	Num     modeltype.Total `json:"num"`
	Time    int64           `json:"time"`
	Wupin   *Wupin          `json:"wupin"`
}

func NewBag(bag *model.Bag) Bag {
	var class *Class
	var wp *Wupin

	if bag.WuPinID <= 0 || bag.WuPin == nil {
		panic("wupin is nil")
	}

	if bag.WuPin.ClassID >= 0 && bag.WuPin.Class != nil && bag.WuPin.Class.Show {
		class = &Class{
			ID:   bag.WuPin.ClassID,
			Name: bag.WuPin.Class.Name,
		}
	} else {
		class = &Class{
			ID:   modeltype.ClassEmptyID,
			Name: modeltype.ClassEmptyName,
		}
	}

	if bag.WuPinID > 0 || bag.WuPin != nil {
		wp = &Wupin{
			ID:        bag.WuPin.ID,
			Name:      bag.WuPin.Name,
			Pic:       bag.WuPin.Pic,
			ClassID:   class.ID,
			ClassOf:   class,
			Tag:       utils.GetSQLNullString(bag.WuPin.Tag),
			HotPrice:  modeltype.GetPrice(bag.WuPin.HotPrice),
			RealPrice: modeltype.GetPrice(bag.WuPin.RealPrice),
			Info:      bag.WuPin.Info,
			Ren:       bag.WuPin.Ren,
			Phone:     bag.WuPin.Phone,
			Email:     utils.GetSQLNullString(bag.WuPin.Email),
			Wechat:    utils.GetSQLNullString(bag.WuPin.WeChat),
			Location:  bag.WuPin.Location,
			BuyTotal:  modeltype.GetTotal(bag.WuPin.BuyTotal),
			BuyDaohuo: modeltype.GetTotal(bag.WuPin.BuyDaoHuo),
			BuyGood:   modeltype.GetTotal(bag.WuPin.BuyGood),
		}
	} else {
		wp = &Wupin{
			ID: modeltype.WupinEmptyID,
		}
	}

	return Bag{
		ID:      bag.ID,
		UserID:  bag.UserID,
		WupinID: bag.WuPinID,
		ClassID: bag.ClassID,
		Num:     bag.Num,
		Time:    bag.Time.Unix(),
		Wupin:   wp,
	}
}

type Data struct {
	List  []Bag `json:"list"`
	Total int   `json:"total"`
}

func NewData(res []model.Bag) Data {
	list := make([]Bag, len(res))
	for _, v := range res {
		if !v.IsBagDown() {
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
