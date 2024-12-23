package config

import "strings"

type FrontConfig struct {
	BasePath    string `yaml:"basepath"`
	TestPath    string `yaml:"testpath"`
	TestPayPath string `yaml:"testpaypath"`
}

func (f *FrontConfig) setDefault() {
	f.BasePath = strings.TrimSpace(f.BasePath)

	if f.BasePath == "" {
		f.BasePath = "/"
	} else {
		if !strings.HasPrefix(f.BasePath, "/") {
			f.BasePath = "/" + f.BasePath
		}

		if strings.HasSuffix(f.BasePath, "/") {
			f.BasePath = strings.TrimRight(f.BasePath, "/")
		}
	}

	if f.TestPayPath == "" {
		f.TestPayPath = "/pay"
	} else {
		if !strings.HasPrefix(f.TestPayPath, "/") {
			f.TestPayPath = "/" + f.TestPayPath
		}

		if strings.HasSuffix(f.TestPayPath, "/") {
			f.TestPayPath = strings.TrimRight(f.TestPayPath, "/")
		}
	}

	if f.TestPath == "" {
		f.TestPath = "/test"
	} else {
		if !strings.HasPrefix(f.TestPath, "/") {
			f.TestPath = "/" + f.TestPath
		}

		if strings.HasSuffix(f.TestPath, "/") {
			f.TestPath = strings.TrimRight(f.TestPath, "/")
		}
	}
}

func (c *FrontConfig) check() ConfigError {
	return nil
}
