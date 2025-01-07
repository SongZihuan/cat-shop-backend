package model

import (
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Key   modeltype.ConfigKeyType          `gorm:"type:varchar(20);not null;uniqueIndex"`
	Value modeltype.ConfigValueType        `gorm:"type:varchar(20);not null"`
	Type  modeltype.TypesOfConfigValueType `gorm:"type:varchar(20);not null"`
}

func NewConfig(Key modeltype.ConfigKeyType, Value modeltype.ConfigValueType) *Config {
	tp, ok := modeltype.ConfigType[Key]
	if !ok {
		panic("bad config key")
	}

	_, ok = modeltype.ConfigInfo[Key]
	if !ok {
		panic("bad config key")
	}

	return &Config{
		Key:   Key,
		Value: Value,
		Type:  tp,
	}
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
