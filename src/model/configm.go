package model

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type ConfigM struct {
	gorm.Model
	Key   modeltype.ConfigKeyType          `gorm:"type:varchar(20);not null;uniqueIndex"`
	Value modeltype.ConfigValueType        `gorm:"type:varchar(20);not null"`
	Type  modeltype.TypesOfConfigValueType `gorm:"type:varchar(20);not null"`
}

func (*ConfigM) TableName() string {
	return "config"
}

func init() {
	if !modelTest[Config, ConfigM]() {
		panic("database error")
	}
}
