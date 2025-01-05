package getwupin

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type Query struct {
	ID uint `form:"id"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Wupin struct {
	ID         uint            `json:"id"`
	Name       string          `json:"name"`
	Pic        string          `json:"pic"`
	ClassID    uint            `json:"classId"`
	ClassOf    *Class          `json:"classOf"`
	Tag        string          `json:"tag,omitempty"`
	HotPrice   int64           `json:"hotPrice,omitempty"`
	RealPrice  int64           `json:"realPrice"`
	Info       string          `json:"info"`
	Ren        string          `json:"ren"`
	Phone      string          `json:"phone"`
	Email      string          `json:"email,omitempty"`
	Wechat     string          `json:"wechat,omitempty"`
	Location   string          `json:"location"`
	BuyMoney   modeltype.Price `json:"buyprice"`
	BuyTotal   modeltype.Total `json:"buytotal"`
	BuyDaohuo  modeltype.Total `json:"buydaohuo"`
	BuyPingJia modeltype.Total `json:"buypingjia"`
	BuyJian    modeltype.Total `json:"buyjiaa"`
	BuyGood    modeltype.Total `json:"buygood"`
}

func NewData(wp *model.Wupin) Wupin {
	var class *Class
	if wp.IsClassDownOrNotShow() {
		class = &Class{
			ID:   modeltype.ClassEmptyID,
			Name: modeltype.ClassEmptyName,
		}
	} else {
		class = &Class{
			ID:   wp.Class.ID,
			Name: wp.Class.Name,
		}
	}

	return Wupin{
		ID:         wp.ID,
		Name:       wp.Name,
		Pic:        wp.Pic,
		ClassID:    class.ID,
		ClassOf:    class,
		Tag:        utils.GetSQLNullString(wp.Tag),
		HotPrice:   modeltype.GetPrice(wp.HotPrice),
		RealPrice:  modeltype.GetPrice(wp.RealPrice),
		Info:       wp.Info,
		Ren:        wp.Ren,
		Phone:      wp.Phone,
		Email:      utils.GetSQLNullString(wp.Email),
		Wechat:     utils.GetSQLNullString(wp.WeChat),
		Location:   wp.Location,
		BuyMoney:   wp.BuyMoney,
		BuyTotal:   wp.BuyTotal,
		BuyDaohuo:  wp.BuyDaoHuo,
		BuyPingJia: wp.BuyPingjia,
		BuyJian:    wp.BuyJian,
		BuyGood:    wp.BuyGood,
	}
}

func NewJsonData(wp *model.Wupin) data.Data {
	return data.NewSuccessWithData(NewData(wp))
}
