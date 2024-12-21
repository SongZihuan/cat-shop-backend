package model

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Image struct {
	gorm.Model
	Type modeltype.ImageType `gorm:"type:uint;not null"`
	Hash string              `gorm:"type:char(64);not null"`
	Time time.Time           `gorm:"type:datetime;not null"`
}

func (img *Image) GetUrl() string {
	return GetImageUrl(img.Type, img.Hash, img.Time.Unix())
}

func GetImageUrl(tp modeltype.ImageType, hash string, time int64) string {
	v := url.Values{}
	tpn, ok := modeltype.ImageTypeToName[tp]
	if !ok {
		return ""
	}

	v.Add("type", tpn)
	v.Add("hash", hash)
	v.Add("time", fmt.Sprintf("%d", time))

	return v.Encode()
}
