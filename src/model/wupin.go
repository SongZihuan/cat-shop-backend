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

	BuyTotal   modeltype.Total `gorm:"type:uint;not null"`
	BuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	BuyPingjia modeltype.Total `gorm:"type:uint;not null"`
	BuyGood    modeltype.Total `gorm:"type:uint;not null"`

	Hot       bool `gorm:"type:boolean;not null;"`
	Show      bool `gorm:"type:boolean;not null;"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func (*WuPin) TableName() string {
	return "wupin"
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
