package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"time"
)

type VideoM struct {
	gorm.Model
	Type modeltype.VideoType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (*VideoM) TableName() string {
	return "video"
}

func init() {
	if !modelTest[Video, VideoM]() {
		panic("database error")
	}
}
