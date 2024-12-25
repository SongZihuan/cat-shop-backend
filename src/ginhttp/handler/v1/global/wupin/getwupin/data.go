package getwupin

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
)

type Query struct {
	ID uint `form:"id"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Wupin struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	ClassID   uint   `json:"classid"`
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

func NewData(wp *model.WuPin) Wupin {
	var class *Class
	if wp.ClassID >= 0 && wp.Class != nil && wp.Class.Show {
		class = &Class{
			ID:   wp.ClassID,
			Name: wp.Class.Name,
		}
	} else {
		class = &Class{
			ID:   0,
			Name: "特殊类别",
		}
	}

	return Wupin{
		ID:        wp.ID,
		Name:      wp.Name,
		Pic:       wp.Pic,
		ClassID:   class.ID,
		ClassOf:   class,
		Tag:       utils.GetSQLNullString(wp.Tag),
		HotPrice:  modeltype.GetPrice(wp.HotPrice),
		RealPrice: modeltype.GetPrice(wp.RealPrice),
		Info:      wp.Info,
		Ren:       wp.Ren,
		Phone:     wp.Phone,
		Email:     utils.GetSQLNullString(wp.Email),
		Wechat:    utils.GetSQLNullString(wp.WeChat),
		Location:  wp.Location,
		BuyTotal:  modeltype.GetTotal(wp.BuyTotal),
		BuyDaohuo: modeltype.GetTotal(wp.BuyDaoHuo),
		BuyGood:   modeltype.GetTotal(wp.BuyGood),
	}
}

func NewJsonData(wp *model.WuPin) data.Data {
	return data.NewSuccessWithData(NewData(wp))
}
