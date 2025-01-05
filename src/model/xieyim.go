package model

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type XieyiM struct {
	gorm.Model
	Type modeltype.XieYiType `gorm:"type:VARCHAR(20);not null"`
	Data string              `gorm:"type:TEXT;not null"`
}

func (*XieyiM) TableName() string {
	return "xieyi"
}

func init() {
	if !modelTest[Xieyi, XieyiM]() {
		panic("database error")
	}
}
