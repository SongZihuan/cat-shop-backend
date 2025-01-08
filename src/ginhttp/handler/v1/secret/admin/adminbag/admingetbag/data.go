package admingetbag

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

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Show bool   `json:"show"`
	Down bool   `json:"down"`
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
	Wupin   Wupin           `json:"wupin"`
	Down    bool            `json:"down"`
}

func NewBag(bag *model.Bag) Bag {
	var class *Class

	if bag.WupinID <= 0 || bag.Wupin == nil {
		panic("wupin is nil")
	}

	if bag.ClassID != bag.Wupin.ClassID {
		panic("class id not equal")
	}

	if bag.ClassID == modeltype.ClassEmptyID {
		class = &Class{
			ID:   modeltype.ClassEmptyID,
			Name: modeltype.ClassEmptyName,
			Show: modeltype.ClassEmptyShow,
			Down: modeltype.ClassEmptyDown,
		}
	} else if bag.ClassID >= 0 && bag.Class != nil {
		class = &Class{
			ID:   bag.Wupin.ClassID,
			Name: bag.Wupin.Class.Name,
			Show: bag.Wupin.Class.Show,
			Down: bag.Wupin.Class.ClassDown,
		}
	} else {
		panic("class is nil")
	}

	return Bag{
		ID:      bag.ID,
		UserID:  bag.UserID,
		WupinID: bag.WupinID,
		ClassID: bag.ClassID,
		Num:     bag.Num,
		Time:    bag.Time.Unix(),
		Wupin: Wupin{
			ID:        bag.Wupin.ID,
			Name:      bag.Wupin.Name,
			Pic:       bag.Wupin.Pic,
			ClassID:   class.ID,
			ClassOf:   class,
			Tag:       utils.GetSQLNullString(bag.Wupin.Tag),
			HotPrice:  bag.Wupin.HotPrice.ToInt64(),
			RealPrice: bag.Wupin.RealPrice.ToInt64(),
			Info:      bag.Wupin.Info,
			Ren:       bag.Wupin.Ren,
			Phone:     bag.Wupin.Phone,
			Email:     utils.GetSQLNullString(bag.Wupin.Email),
			Wechat:    utils.GetSQLNullString(bag.Wupin.WeChat),
			Location:  bag.Wupin.Location,
			BuyTotal:  bag.Wupin.BuyTotal.ToInt64(),
			BuyDaohuo: bag.Wupin.BuyDaoHuo.ToInt64(),
			BuyGood:   bag.Wupin.BuyGood.ToInt64(),
		},
		Down: bag.IsBagDown(),
	}
}

type Data struct {
	List     []Bag `json:"list"`
	Total    int   `json:"total"`
	MaxCount int   `json:"maxpage"`
}

func NewData(res []model.Bag, maxcount int) Data {
	list := make([]Bag, len(res))
	for _, v := range res {
		list = append(list, NewBag(&v))
	}

	return Data{
		List:     list,
		Total:    len(res),
		MaxCount: maxcount,
	}
}

func NewJsonData(res []model.Bag, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(res, maxcount))
}
