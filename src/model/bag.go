package model

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type Bag struct {
	gorm.Model
	UserID    uint            `gorm:"not null"`
	User      *User           `gorm:"foreignKey:UserID"`
	WupinID   uint            `gorm:"not null"`
	Wupin     *Wupin          `gorm:"foreignKey:WupinID"`
	ClassID   uint            `gorm:"not null"`
	Class     *Class          `gorm:"foreignKey:ClassID"`
	Num       modeltype.Total `gorm:"type:uint;not null"`
	Time      time.Time       `gorm:"type:datetime;not null"`
	WupinDown bool            `gorm:"type:boolean;not null"`
	ClassShow bool            `gorm:"type:boolean;not null;"`
	ClassDown bool            `gorm:"type:boolean;not null;"`
}

func NewBag(user *User, wupin *Wupin, num modeltype.Total) *Bag {
	tmp1 := *wupin
	tmp2 := *wupin.Class
	tmp3 := *user

	return &Bag{
		UserID:    user.ID,
		User:      &tmp3,
		WupinID:   wupin.ID,
		Wupin:     &tmp1,
		ClassID:   wupin.ClassID,
		Class:     &tmp2,
		Num:       num,
		Time:      time.Now(),
		WupinDown: wupin.IsWupinDown(),
		ClassShow: wupin.Class.IsClassShow(),
		ClassDown: wupin.Class.IsClassDown(),
	}
}

func (*Bag) TableName() string {
	return "bag"
}

func (bag *Bag) Add(num int) bool {
	bag.Num += modeltype.Total(num)
	if bag.Num <= 0 {
		bag.Num = 0
		return false
	}
	return true
}

func (bag *Bag) isClassDown() bool {
	if bag.Class == nil {
		return bag.ClassDown
	} else {
		if bag.Class.ID != bag.ClassID {
			panic("class id not equal")
		}

		return bag.Class.IsClassDown()
	}
}

func (bag *Bag) isWupinDown() bool {
	if bag.Wupin == nil {
		return bag.WupinDown || bag.ClassShow
	} else {
		if bag.WupinID != bag.Wupin.ID {
			panic("wupin id not equal")
		}

		return bag.Wupin.IsWupinDown()
	}
}

func (bag *Bag) IsBagDown() bool {
	return bag.isWupinDown() || bag.isClassDown()
}

func (bag *Bag) IsBagShow() bool {
	return bag.Num > 0
}

func (bag *Bag) IsBagNotShow() bool {
	return !bag.IsBagShow()
}

func (bag *Bag) IsBagCanSale() bool {
	return bag.IsBagShow() && !bag.IsBagDown()
}

func (bag *Bag) IsBagCanNotSale() bool {
	return !bag.IsBagCanNotSale()
}
