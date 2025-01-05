package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

type Xieyi struct {
	gorm.Model
	Type modeltype.XieYiType `gorm:"type:VARCHAR(20);not null"`
	Data string              `gorm:"type:TEXT;not null"`
}

func (*Xieyi) TableName() string {
	return "xieyi"
}

func NewXieyi(xieyiType modeltype.XieYiType, content string) *Xieyi {
	if xieyiType == "" {
		xieyiType = modeltype.XieYiDefault
	}

	return &Xieyi{
		Type: modeltype.XieyiUser,
		Data: content,
	}
}
