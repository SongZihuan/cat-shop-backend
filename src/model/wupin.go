package model

import (
	"database/sql"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

type WuPin struct {
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

	BuyMoney   modeltype.Price `gorm:"type:uint;not null"`
	BuyTotal   modeltype.Total `gorm:"type:uint;not null"`
	BuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	BuyPingjia modeltype.Total `gorm:"type:uint;not null"`
	BuyGood    modeltype.Total `gorm:"type:uint;not null"`
	BuyJian    modeltype.Total `gorm:"type:uint;not null"`

	Hot       bool `gorm:"type:boolean;not null;"`
	Show      bool `gorm:"type:boolean;not null;"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func (*WuPin) TableName() string {
	return "wupin"
}

func (w *WuPin) BuyNow(r *BuyRecord) bool {
	if r.WuPinID != w.ID || r.WuPin == nil || r.WuPin.ID != w.ID {
		return false
	}

	w.BuyMoney += r.TotalPrice
	w.BuyTotal += 1
	w.BuyJian += r.Num
	return true
}

func (w *WuPin) BackPayNow(r *BuyRecord) bool {
	if r.WuPinID != w.ID || r.WuPin == nil || r.WuPin.ID != w.ID {
		return false
	}

	w.BuyMoney -= r.TotalPrice
	w.BuyTotal -= 1
	w.BuyJian -= r.Num

	if w.BuyMoney <= 0 {
		w.BuyMoney = 0
	}

	if w.BuyTotal < 0 {
		w.BuyTotal = 0
	}

	if w.BuyJian < 0 {
		w.BuyJian = 0
	}

	return true
}

func (w *WuPin) Daohuo(r *BuyRecord) bool {
	if r.WuPinID != w.ID || r.WuPin == nil || r.WuPin.ID != w.ID {
		return true
	}
	w.BuyDaoHuo += 1
	return false
}

func (w *WuPin) PingJia(r *BuyRecord, isGood bool) bool {
	if r.WuPinID != w.ID || r.WuPin == nil || r.WuPin.ID != w.ID {
		return true
	}

	w.BuyPingjia += 1

	if isGood {
		w.BuyGood += 1
	}

	return false
}

func (w *WuPin) GetRealPrice() modeltype.Price {
	if w.RealPrice >= 0 {
		return w.RealPrice
	}
	return 0
}

func (w *WuPin) GetPrice() modeltype.Price {
	realPrice := w.GetRealPrice()

	if !w.HotPrice.Valid || w.HotPrice.V < 0 {
		return realPrice
	}

	if w.HotPrice.V < realPrice {
		return w.HotPrice.V
	}

	return realPrice
}

func (w *WuPin) GetFacePrice() modeltype.Price {
	return w.GetPrice()
}

func (w *WuPin) GetPriceTotal(num modeltype.Total) modeltype.Price {
	price := w.GetPrice()
	if price < 0 {
		return 0
	}
	return modeltype.Price(int64(price) * int64(num))
}

func (w *WuPin) IsClassDownOrNotShow() bool {
	if w.Class == nil {
		if w.ClassDown {
			return true // 下架状态 均返回fakse
		} else {
			if w.ClassShow {
				return false // 非下架状态，Show为true表示展示
			} else {
				return true // 非下架状态，Show为false表示隐藏
			}
		}
	} else {
		if w.ClassID != w.Class.ID {
			panic("class id not equal")
		}

		return w.Class.IsClassDownOrNotShow()
	}
}

func (w *WuPin) IsClassDown() bool {
	if w.Class == nil {
		return w.ClassDown
	} else {
		if w.ClassID != w.Class.ID {
			panic("class id not equal")
		}

		return w.Class.IsClassDown()
	}
}

func (w *WuPin) IsWupinDown() bool {
	if w.Class == nil {
		return !w.Show || w.ClassDown
	} else {
		if w.ClassID != w.Class.ID {
			panic("class id not equal")
		}

		return !w.Show || w.Class.IsClassDown()
	}
}
