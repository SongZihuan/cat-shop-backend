package model

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"gorm.io/gorm"
	"net/url"
	"os"
	"path"
	"time"
)

type Video struct {
	gorm.Model
	Type modeltype.VideoType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (*Video) TableName() string {
	return "video"
}

func NewVideo(tp modeltype.VideoType, file []byte) (*Video, error) {
	vid := &Video{
		Type: tp,
		Hash: utils.SHA256(file),
		Time: time.Now(),
	}

	err := vid.saveFile(file)
	if err != nil {
		return nil, err
	}
	return vid, nil
}

func (vid *Video) saveFile(file []byte) error {
	return os.WriteFile(vid.SavePath(), file, os.ModePerm)
}

func (vid *Video) SavePath() string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	hash := vid.Hash
	if len(hash) != 64 {
		return ""
	}

	basePath, ok := config.Config().File.Video[vid.Type]
	if !ok {
		return ""
	}

	return path.Join(basePath, fmt.Sprintf("%d", vid.Time.Unix()), fmt.Sprintf("%s.dat", hash))
}

func (vid *Video) GetUrl() string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	return config.Config().Yaml.Http.ResourceBaseAPI + modeltype.VideoPath + "?" + vid.GetQuery()
}

func (vid *Video) GetQuery() string {
	v := url.Values{}
	tpn, ok := modeltype.VideoTypeToName[vid.Type]
	if !ok {
		return ""
	}

	v.Add("type", tpn)
	v.Add("hash", vid.Hash)
	v.Add("time", fmt.Sprintf("%d", vid.Time.Unix()))

	return v.Encode()
}
