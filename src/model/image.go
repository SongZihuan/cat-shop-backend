package model

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/loadpath"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"gorm.io/gorm"
	"net/url"
	"os"
	"path"
	"time"
)

type Image struct {
	gorm.Model
	Type modeltype.ImageType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (*Image) TableName() string {
	return "image"
}

func NewImage(tp modeltype.ImageType, file []byte) (*Image, error) {
	img := &Image{
		Type: tp,
		Hash: utils.SHA256(file),
		Time: time.Now(),
	}

	err := img.saveFile(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (img *Image) saveFile(file []byte) error {
	return os.WriteFile(img.SavePath(), file, os.ModePerm)
}

func (img *Image) SavePath() string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	hash := img.Hash
	if len(hash) != 64 {
		return ""
	}

	basePath, ok := config.Config().File.Image[img.Type]
	if !ok {
		return ""
	}

	return path.Join(basePath, fmt.Sprintf("%d", img.Time.Unix()), fmt.Sprintf("%s.dat", hash))
}

func (img *Image) GetUrl() string {
	return loadpath.GetImagePath() + "?" + img.GetQuery()
}

func (img *Image) GetQuery() string {
	v := url.Values{}
	tpn, ok := modeltype.ImageTypeToName[img.Type]
	if !ok {
		return ""
	}

	v.Add("type", tpn)
	v.Add("hash", img.Hash)
	v.Add("time", fmt.Sprintf("%d", img.Time.Unix()))

	return v.Encode()
}
