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
	WupinHot  bool            `gorm:"type:boolean;not null"`
	WupinDown bool            `gorm:"type:boolean;not null"`
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

func (bag *Bag) IsClassDown() bool {
	if bag.Class == nil {
		return bag.ClassDown
	} else {
		if bag.Class.ID != bag.ClassID {
			panic("class id not equal")
		}

		return bag.Class.IsClassDown()
	}
}

func (bag *Bag) IsWupinDown() bool {
	if bag.WuPin == nil {
		return bag.WupinDown || bag.ClassShow
	} else {
		if bag.WuPinID != bag.WuPin.ID {
			panic("wupin id not equal")
		}

		return bag.WuPin.IsWupinDown()
	}
}

func (bag *Bag) IsBagDown() bool {
	return bag.IsWupinDown() || bag.IsClassDown()
}

func (bag *Bag) IsBagShow() bool {
	return bag.Num > 0 && !bag.IsBagDown()
}
