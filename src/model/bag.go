package model

import (
	"gorm.io/gorm"
	"time"
)

type Bag struct {
	gorm.Model
	UserID  uint      `gorm:"not null"`
	User    *User     `gorm:"foreignKey:UserID"`
	WuPinID uint      `gorm:"not null"`
	WuPin   *WuPin    `gorm:"foreignKey:WuPinID"`
	ClassID uint      `gorm:"not null"`
	Class   *Class    `gorm:"foreignKey:ClassID"`
	Num     Total     `gorm:"type:uint;not null"`
	Time    time.Time `gorm:"type:datetime;not null"`
}
