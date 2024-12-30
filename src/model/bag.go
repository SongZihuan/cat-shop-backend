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

func (bag *Bag) IsClassDownOrNotShow() bool {
	if bag.WuPin == nil {
		if bag.ClassDown {
			return true // 下架状态 均返回fakse
		} else {
			if bag.ClassShow {
				return false // 非下架状态，Show为true表示展示
			} else {
				return true // 非下架状态，Show为false表示隐藏
			}
		}
	} else {
		if bag.WuPin.ID != bag.WuPinID {
			panic("wupin id not equal")
		}

		return bag.WuPin.IsClassDownOrNotShow()
	}
}

func (bag *Bag) IsClassDown() bool {
	if bag.WuPin == nil {
		return bag.ClassDown
	} else {
		if bag.WuPin.ID != bag.WuPinID {
			panic("wupin id not equal")
		}

		return bag.WuPin.IsClassDown()
	}
}

func (bag *Bag) IsWupinDown() bool {
	if bag.WuPin == nil {
		return !bag.WuPinShow || bag.ClassShow
	} else {
		if bag.WuPinID != bag.WuPin.ID {
			panic("wupin id not equal")
		}

		return !bag.WuPin.IsWupinDown()
	}
}

func (bag *Bag) IsBagDownNotBecauseNum() bool {
	return bag.IsWupinDown()
}

func (bag *Bag) IsBagDown() bool {
	if bag.Num <= 0 {
		return true
	}

	return bag.IsBagDownNotBecauseNum()
}
