package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Key   modeltype.ConfigKeyType          `gorm:"type:varchar(20);not null;uniqueIndex"`
	Value modeltype.ConfigValueType        `gorm:"type:varchar(20);not null"`
	Type  modeltype.TypesOfConfigValueType `gorm:"type:varchar(20);not null"`
}

func (c *Config) Default() modeltype.ConfigValueType {
	if !config.IsReady() {
		panic("config is not ready")
	}

	if c.Key == modeltype.KeyPasswordFrontHash {
		return modeltype.ConfigValueType(config.Config().Yaml.Password.Front)
	}

	return ""
}

func (c *Config) GetValue() modeltype.ConfigValueType {
	if c.Value == "" {
		return c.Default()
	}
	return c.Value
}
