package config

type FrontConfig struct {
	BasePath    string `yaml:"basepath"`
	TestPath    string `yaml:"testpath"`
	TestPayPath string `yaml:"testpaypath"`
}

func (f *FrontConfig) setDefault() {
	f.BasePath = processURL(f.BasePath)
	f.TestPath = processURL(f.TestPath, "/test")
	f.TestPayPath = processURL(f.TestPayPath, "/pay")
}

func (c *FrontConfig) check() ConfigError {
	return nil
}
