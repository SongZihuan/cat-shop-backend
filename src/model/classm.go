package model

import (
	"gorm.io/gorm"
)

type ClassM struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Show      bool   `gorm:"type:boolean;not null"`
	ClassDown bool   `gorm:"type:boolean;not null"` // 下架所有商品
}

func (*ClassM) TableName() string {
	return "class"
}

func init() {
	if !modelTest[Class, ClassM]() {
		panic("database error")
	}
}
