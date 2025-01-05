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

	BuyMoney   modeltype.Price `gorm:"type:uint;not null"`
	BuyTotal   modeltype.Total `gorm:"type:uint;not null"`
	BuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	BuyPingjia modeltype.Total `gorm:"type:uint;not null"`
	BuyGood    modeltype.Total `gorm:"type:uint;not null"`
	BuyJian    modeltype.Total `gorm:"type:uint;not null"`

	Hot       bool `gorm:"type:boolean;not null;"`
	Down      bool `gorm:"type:boolean;not null;"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func (*Wupin) TableName() string {
	return "wupin"
}

func (w *Wupin) BuyNow(r *BuyRecord) bool {
	if r.WupinID != w.ID || r.Wupin == nil || r.Wupin.ID != w.ID {
		return false
	}

	w.BuyMoney += r.TotalPrice
	w.BuyTotal += 1
	w.BuyJian += r.Num
	return true
}

func (w *Wupin) BackPayNow(r *BuyRecord) bool {
	if r.WupinID != w.ID || r.Wupin == nil || r.Wupin.ID != w.ID {
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

func (w *Wupin) Daohuo(r *BuyRecord) bool {
	if r.WupinID != w.ID || r.Wupin == nil || r.Wupin.ID != w.ID {
		return true
	}
	w.BuyDaoHuo += 1
	return false
}

func (w *Wupin) PingJia(r *BuyRecord, isGood bool) bool {
	if r.WupinID != w.ID || r.Wupin == nil || r.Wupin.ID != w.ID {
		return true
	}

	w.BuyPingjia += 1

	if isGood {
		w.BuyGood += 1
	}

	return false
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

func (w *Wupin) IsClassDownOrNotShow() bool {
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

func (w *Wupin) IsClassDown() bool {
	if w.Class == nil {
		return w.ClassDown
	} else {
		if w.ClassID != w.Class.ID {
			panic("class id not equal")
		}

		return w.Class.IsClassDown()
	}
}

func (w *Wupin) IsWupinDown() bool {
	if w.Class == nil && w.ClassID != w.Class.ID {
		panic("class id not equal")
	}

	return w.Down || w.ClassDown || (w.Class != nil && w.Class.IsClassDown())
}
