package config

import (
	"github.com/SuperH-0630/cat-shop-back/src/flagparser"
)

func newConfig() ConfigStruct {
	return ConfigStruct{
		configReady:   false,
		yamlHasParser: false,
	}
}

func InitConfig() ConfigError {
	if !flagparser.IsReady() {
		return NewConfigError("flag not ready")
	}

	config = newConfig()
	err := config.ready()
	if err != nil && err.IsError() {
		return err
	}

	if !config.configReady {
		return NewConfigError("config not ready")
	}

	return nil
}

func IsReady() bool {
	return config.yamlHasParser && config.configReady
}

func Config() *ConfigStruct {
	if !IsReady() {
		panic("config not ready")
	}
	return &config
}

var config ConfigStruct
