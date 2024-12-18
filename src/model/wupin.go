package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type WuPin struct {
	gorm.Model
	Name    string         `gorm:"type:varchar(20);not null"`
	ClassID uint           `gorm:"not null"`
	Class   *Class         `gorm:"foreignKey:ClassID"`
	Tag     sql.NullString `gorm:"type:varchar(20)"`

	HotPrice  PriceNull `gorm:"type:uint;"`
	RealPrice Price     `gorm:"type:uint;not null;"`

	Info     string         `gorm:"type:text;not null"`
	Ren      string         `gorm:"type:varchar(20);not null"`
	Phone    string         `gorm:"type:varchar(30);not null"`
	WeChat   sql.NullString `gorm:"type:varchar(50);"`
	Email    sql.NullString `gorm:"type:varchar(50);"`
	Location string         `gorm:"type:varchar(200);not null"`

	BuyTotal  Total `gorm:"type:uint;not null"`
	BuyDaoHuo Total `gorm:"type:uint;not null"`
	BuyGood   Total `gorm:"type:uint;not null"`

	IsHot bool `gorm:"type:boolean;not null;"`
}
