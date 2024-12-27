package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type Bag struct {
	gorm.Model
	UserID    uint            `gorm:"not null"`
	User      *User           `gorm:"foreignKey:UserID"`
	WuPinID   uint            `gorm:"not null"`
	WuPin     *WuPin          `gorm:"foreignKey:WuPinID"`
	ClassID   uint            `gorm:"not null"`
	Class     *Class          `gorm:"foreignKey:ClassID"`
	Num       modeltype.Total `gorm:"type:uint;not null"`
	Time      time.Time       `gorm:"type:datetime;not null"`
	WuPinShow bool            `gorm:"type:boolean;not null"`
	WupinHot  bool            `gorm:"type:boolean;not null"`
	ClassShow bool            `gorm:"type:boolean;not null;"`
	ClassDown bool            `gorm:"type:boolean;not null;"`
}

func (*Bag) TableName() string {
	return "bag"
}

func (bag *Bag) Add(num int) bool {
	bag.Num += modeltype.Total(num)
	return bag.Num != 0
}
