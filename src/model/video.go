package model

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Video struct {
	gorm.Model
	Type modeltype.VideoType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (vid *Video) GetUrl() string {
	return GetVideoUrl(vid.Type, vid.Hash, vid.Time.Unix())
}

func GetVideoUrl(tp modeltype.VideoType, hash string, time int64) string {
	v := url.Values{}
	tpn, ok := modeltype.VideoTypeToName[tp]
	if !ok {
		return ""
	}

	v.Add("type", tpn)
	v.Add("hash", hash)
	v.Add("time", fmt.Sprintf("%d", time))

	return v.Encode()
}
