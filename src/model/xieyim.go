package model

import "gorm.io/gorm"

type XieyiM struct {
	gorm.Model
	Data string `gorm:"type:TEXT;not null"`
}

func (*XieyiM) TableName() string {
	return "xieyi"
}

func init() {
	if !modelTest[Xieyi, XieyiM]() {
		panic("database error")
	}
}
