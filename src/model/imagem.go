package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type ImageM struct {
	gorm.Model
	Type modeltype.ImageType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (*ImageM) TableName() string {
	return "image"
}

func init() {
	if !modelTest[Image, ImageM]() {
		panic("database error")
	}
}
