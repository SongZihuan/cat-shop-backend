package model

import (
	"database/sql"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type Wupin struct {
	gorm.Model
	Name    string         `gorm:"type:varchar(20);not null"`
	Pic     string         `gorm:"type:varchar(150);not null"`
	ClassID uint           `gorm:"not null"`
	Class   *Class         `gorm:"foreignKey:ClassID"`
	Tag     sql.NullString `gorm:"type:varchar(20)"`

	HotPrice  modeltype.PriceNull `gorm:"type:uint;"`
	RealPrice modeltype.Price     `gorm:"type:uint;not null;"`

	Info     string         `gorm:"type:text;not null"`
	Ren      string         `gorm:"type:varchar(20);not null"`
	Phone    string         `gorm:"type:varchar(30);not null"`
	WeChat   sql.NullString `gorm:"type:varchar(50);"`
	Email    sql.NullString `gorm:"type:varchar(50);"`
	Location string         `gorm:"type:varchar(200);not null"`

	BuyPrice   modeltype.Price `gorm:"type:uint;not null"` // 购物总金额
	BuyTotal   modeltype.Total `gorm:"type:uint;not null"` // 购物总人数
	BuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	BuyPingjia modeltype.Total `gorm:"type:uint;not null"`
	BuyGood    modeltype.Total `gorm:"type:uint;not null"`
	BuyJian    modeltype.Total `gorm:"type:uint;not null"`

	Hot       bool `gorm:"type:boolean;not null;"`
	WupinDown bool `gorm:"type:boolean;not null;"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func (*Wupin) TableName() string {
	return "wupin"
}

func NewWupin(name string, pic string, class *Class, tag string, hotPrice modeltype.PriceNull, realPrice modeltype.Price, info string, ren string, phone string, email string, wechat string, location string, hot bool, down bool) *Wupin {
	tmp := *class

	return &Wupin{
		Name:      name,
		Pic:       pic,
		ClassID:   class.ID,
		Class:     &tmp,
		Tag:       sql.NullString{String: tag, Valid: len(tag) != 0},
		HotPrice:  hotPrice,
		RealPrice: realPrice,
		Info:      info,
		Ren:       ren,
		Phone:     phone,
		WeChat:    sql.NullString{String: wechat, Valid: len(wechat) != 0},
		Email:     sql.NullString{String: email, Valid: len(email) != 0},
		Location:  location,
		Hot:       hot,
		WupinDown: down,
		ClassShow: class.IsClassShow(),
		ClassDown: class.IsClassDown(),
	}
}

func (w *Wupin) BuyNow(r *BuyRecord) bool {
	w.BuyPrice += r.TotalPrice
	w.BuyTotal += 1
	w.BuyJian += r.Num
	return true
}

func (w *Wupin) BuyQuXiao(r *BuyRecord) bool {
	w.BuyPrice -= r.TotalPrice
	w.BuyTotal -= 1
	w.BuyJian -= r.Num
	return true
}

func (w *Wupin) TuiHuoBeforeFaHuo(r *BuyRecord) bool {
	w.BuyPrice -= r.TotalPrice
	w.BuyTotal -= 1
	w.BuyJian -= r.Num

	if w.BuyPrice <= 0 {
		w.BuyPrice = 0
	}

	if w.BuyTotal < 0 {
		w.BuyTotal = 0
	}

	if w.BuyJian < 0 {
		w.BuyJian = 0
	}

	return true
}

func (w *Wupin) Daohuo() bool {
	w.BuyDaoHuo += 1
	return false
}

func (w *Wupin) PingJia(isGood bool) bool {
	w.BuyPingjia += 1
	if isGood {
		w.BuyGood += 1
	}

	return false
}

func (w *Wupin) TuiHuoAfterFaHuo(r *BuyRecord) bool {
	w.BuyPrice -= r.TotalPrice
	w.BuyTotal -= 1
	w.BuyJian -= r.Num
	w.BuyDaoHuo -= 1

	if r.IsGood.Valid {
		w.BuyPingjia -= 1
		if r.IsGood.Bool {
			w.BuyGood -= 1
		}
	}

	if w.BuyPrice <= 0 {
		w.BuyPrice = 0
	}

	if w.BuyTotal <= 0 {
		w.BuyTotal = 0
	}

	if w.BuyJian <= 0 {
		w.BuyJian = 0
	}

	if w.BuyDaoHuo <= 0 {
		w.BuyDaoHuo = 0
	}

	if w.BuyPingjia <= 0 {
		w.BuyPingjia = 0
	}

	if w.BuyGood <= 0 {
		w.BuyGood = 0
	}
	return true
}

func (w *Wupin) GetRealPrice() modeltype.Price {
	if w.RealPrice >= 0 {
		return w.RealPrice
	}
	return 0
}

func (w *Wupin) GetPrice() modeltype.Price {
	realPrice := w.GetRealPrice()

	if !w.HotPrice.Valid || w.HotPrice.V < 0 {
		return realPrice
	}

	if w.HotPrice.V < realPrice {
		return w.HotPrice.V
	}

	return realPrice
}

func (w *Wupin) GetFacePrice() modeltype.Price {
	return w.GetPrice()
}

func (w *Wupin) GetPriceTotal(num modeltype.Total) modeltype.Price {
	price := w.GetPrice()
	if price < 0 {
		return 0
	}
	return modeltype.Price(int64(price) * int64(num))
}

func (w *Wupin) isClassDown() bool {
	if w.Class == nil {
		return w.ClassDown
	} else {
		if w.ClassID != w.Class.ID {
			panic("class id not equal")
		}

		return w.Class.IsClassDown()
	}
}

func (w *Wupin) isWupinDown() bool {
	if w.Class == nil && w.ClassID != w.Class.ID {
		panic("class id not equal")
	}

	//nolint
	if w.ClassID == modeltype.ClassEmptyID && !modeltype.ClassEmptyDown {
		return true
	}

	return w.WupinDown || w.ClassDown || (w.Class != nil && w.Class.IsClassDown())
}

func (w *Wupin) IsWupinDown() bool {
	return w.IsWupinHot() || w.isClassDown()
}

func (m *Wupin) IsWupinShow() bool {
	return !m.IsWupinDown()
}

func (m *Wupin) IsWupinHot() bool {
	return m.IsWupinShow() && m.Hot
}

func (w *Wupin) UpdateNormalInfo(name string, pic string, class *Class, tag string, hotPrice modeltype.PriceNull, realPrice modeltype.Price, info string, hot bool, down bool) bool {
	tmp := *class
	oldDown := w.IsWupinDown()

	w.Name = name
	w.ClassID = class.ID
	w.Class = &tmp
	w.Tag = sql.NullString{String: tag, Valid: len(tag) != 0}
	w.HotPrice = hotPrice
	w.RealPrice = realPrice
	w.Info = info
	w.Hot = hot
	w.WupinDown = down
	w.ClassShow = class.IsClassShow()
	w.ClassDown = class.IsClassDown()

	if len(w.Pic) == 0 || len(pic) != 0 {
		w.Pic = pic
	}

	return oldDown != w.IsWupinDown()
}

func (w *Wupin) UpdateShopInfo(ren string, phone string, email string, wechat string, location string) bool {
	w.Ren = ren
	w.Phone = phone
	w.Email = sql.NullString{String: email, Valid: len(email) != 0}
	w.WeChat = sql.NullString{String: wechat, Valid: len(wechat) != 0}
	w.Location = location

	return true
}

func (w *Wupin) UpdateInfo(name string, pic string, class *Class, tag string, hotPrice modeltype.PriceNull, realPrice modeltype.Price, info string, ren string, phone string, email string, wechat string, location string, hot bool, down bool) bool {
	res1 := w.UpdateNormalInfo(name, pic, class, tag, hotPrice, realPrice, info, hot, down)
	res2 := w.UpdateShopInfo(ren, phone, email, wechat, location)
	return res1 || res2
}

func (w *Wupin) IsWupinCanSale() bool {
	return !w.IsWupinDown()
}

func (w *Wupin) IsWupinCanNotSale() bool {
	return !w.IsWupinCanSale()
}
