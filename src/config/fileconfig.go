package config

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"path"
)

type FileConfig struct {
	LocalPath string `yaml:"localpath"`
}

func (f *FileConfig) setDefault() {
	if f.LocalPath == "" {
		f.LocalPath = "."
	}
}

func (f *FileConfig) check() ConfigError {
	for _, name := range model.ImageTypeToName {
		p := path.Join(f.LocalPath, name)
		if utils.IsExists(p) {
			if !utils.IsDir(p) {
				return NewConfigError(fmt.Sprintf("%s is not a directory", p))
			}
		} else {
			err := utils.MakeDir(p)
			if err != nil {
				return NewConfigError(fmt.Sprintf("create directory %s error: %s", p, err.Error()))
			}
		}
	}

	return nil
}
