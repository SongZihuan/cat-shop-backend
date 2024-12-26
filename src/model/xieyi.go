package model

import "gorm.io/gorm"

type Xieyi struct {
	gorm.Model
	Data string `gorm:"type:TEXT;not null"`
}

func (*Xieyi) TableName() string {
	return "xieyi"
}
