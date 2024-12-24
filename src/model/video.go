package model

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/config"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"net/url"
	"time"
)

const VideoPath = "/v1/video"

type Video struct {
	gorm.Model
	Type modeltype.VideoType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (vid *Video) GetUrl() string {
	if !config.IsReady() {
		panic("config is not ready")
	}

	return config.Config().Yaml.Http.ResourceBaseAPI + VideoPath + "?" + vid.GetQuery()
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
