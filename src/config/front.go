package config

import "github.com/SongZihuan/cat-shop-backend/src/utils"

type FrontConfig struct {
	BasePath    string `yaml:"basepath"`
	TestPath    string `yaml:"testpath"`
	TestPayPath string `yaml:"testpaypath"`
}

func (f *FrontConfig) setDefault() {
	f.BasePath = utils.ProcessPath(f.BasePath)
	f.TestPath = utils.ProcessPath(f.TestPath, "/test")
	f.TestPayPath = utils.ProcessPath(f.TestPayPath, "/pay")
}

func (c *FrontConfig) check() ConfigError {
	if !utils.IsValidURLPath(c.BasePath) {
		return NewConfigError("front base path is not valid")
	}

	if !utils.IsValidURLPath(c.TestPath) {
		return NewConfigError("front test path is not valid")
	}

	if !utils.IsValidURLPath(c.TestPayPath) {
		return NewConfigError("front test pay path is not valid")
	}

	if !utils.IsValidURLPath(c.BasePath + c.TestPath) {
		return NewConfigError("front test path is not valid")
	}

	if !utils.IsValidURLPath(c.BasePath + c.TestPath + c.TestPayPath) {
		return NewConfigError("front test pay path is not valid")
	}

	return nil
}
