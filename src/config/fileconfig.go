package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
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

func (f *FileConfig) check(fl *FileLocationConfig) ConfigError {
	for tp, name := range modeltype.ImageTypeToName {
		p := path.Join(f.LocalPath, "image", name)
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

		fl.Image[tp] = p
	}

	for tp, name := range modeltype.VideoTypeToName {
		p := path.Join(f.LocalPath, "image", name)
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

		fl.Video[tp] = p
	}

	return nil
}
