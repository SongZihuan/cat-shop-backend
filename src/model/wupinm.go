package model

import (
	"database/sql"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type WupinM struct {
	gorm.Model
	Name    string         `gorm:"type:varchar(20);not null"`
	Pic     string         `gorm:"type:varchar(150);not null"`
	ClassID uint           `gorm:"not null"`
	Tag     sql.NullString `gorm:"type:varchar(20)"`

	HotPrice  modeltype.PriceNull `gorm:"type:uint;"`
	RealPrice modeltype.Price     `gorm:"type:uint;not null;"`

	Info     string         `gorm:"type:text;not null"`
	Ren      string         `gorm:"type:varchar(20);not null"`
	Phone    string         `gorm:"type:varchar(30);not null"`
	WeChat   sql.NullString `gorm:"type:varchar(50);"`
	Email    sql.NullString `gorm:"type:varchar(50);"`
	Location string         `gorm:"type:varchar(200);not null"`

	BuyPrice modeltype.Price `gorm:"type:uint;not null"`
	BuyTotal modeltype.Total `gorm:"type:uint;not null"`

	BuyDaoHuo  modeltype.Total `gorm:"type:uint;not null"`
	BuyPingjia modeltype.Total `gorm:"type:uint;not null"`
	BuyGood    modeltype.Total `gorm:"type:uint;not null"`
	BuyJian    modeltype.Total `gorm:"type:uint;not null"`

	Hot       bool `gorm:"type:boolean;not null;"`
	WupinDown bool `gorm:"type:boolean;not null;"`
	ClassShow bool `gorm:"type:boolean;not null;"`
	ClassDown bool `gorm:"type:boolean;not null;"`
}

func (*WupinM) TableName() string {
	return "wupin"
}

func init() {
	if !modelTest[Wupin, WupinM]() {
		panic("database error")
	}
}
