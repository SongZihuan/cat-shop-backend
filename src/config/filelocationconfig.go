package config

import "github.com/SongZihuan/cat-shop-backend/src/model/modeltype"

const DefaultFileMapSize = 10

type FileLocationConfig struct {
	Image map[modeltype.ImageType]string
	Video map[modeltype.VideoType]string
}

func (f *FileLocationConfig) init() error {
	f.Image = make(map[modeltype.ImageType]string, DefaultFileMapSize)
	f.Video = make(map[modeltype.VideoType]string, DefaultFileMapSize)
	return nil
}
