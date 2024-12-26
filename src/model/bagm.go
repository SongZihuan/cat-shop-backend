package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type BagM struct {
	gorm.Model
	UserID    uint            `gorm:"not null"`
	WuPinID   uint            `gorm:"not null"`
	WuPinShow bool            `gorm:"not null"`
	ClassID   uint            `gorm:"not null"`
	Num       modeltype.Total `gorm:"type:uint;not null"`
	Time      time.Time       `gorm:"type:datetime;not null"`
}

func (*BagM) TableName() string {
	return "bag"
}

func init() {
	if !modelTest[Bag, BagM]() {
		panic("database error")
	}
}
