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
	return nil
}
