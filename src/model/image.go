package model

import (
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type ImageType int

const (
	XieYiImage  ImageType = 1
	WuPinImage  ImageType = 2
	ConfigImage ImageType = 3
	AvatarImage ImageType = 4
)

var NameToImageType = map[string]ImageType{
	"XieYi":  XieYiImage,
	"WuPin":  WuPinImage,
	"Config": ConfigImage,
	"Avatar": AvatarImage,
}

var ImageTypeToName = map[ImageType]string{
	XieYiImage:  "XieYi",
	WuPinImage:  "WuPin",
	ConfigImage: "Config",
	AvatarImage: "Avatar",
}

type Image struct {
	gorm.Model
	Type ImageType `gorm:"type:uint;not null"`
	Hash string    `gorm:"type:char(64);not null"`
	Time time.Time `gorm:"type:datetime;not null"`
}

func (img *Image) getUrl() string {
	v := url.Values{}
	v.Add("type", ImageTypeToName[img.Type])
	v.Add("hash", img.Hash)
	v.Add("time", fmt.Sprintf("%d", img.Time.Unix()))

	return v.Encode()
}
