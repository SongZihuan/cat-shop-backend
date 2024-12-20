package model

import (
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"time"
)

type VideoType int

const (
	XieYiVide VideoType = 1
	WuPinVide VideoType = 2
)

var NameToVideoType = map[string]VideoType{
	"XieYi": XieYiVide,
	"WuPin": WuPinVide,
}

var VideoTypeToName = map[VideoType]string{
	XieYiVide: "XieYi",
	WuPinVide: "WuPin",
}

type Video struct {
	gorm.Model
	Type VideoType `gorm:"type:uint;not null"`
	Hash string    `gorm:"type:char(64);not null"`
	Time time.Time `gorm:"type:datetime;not null"`
}

func (vid *Video) getUrl() string {
	v := url.Values{}
	v.Add("type", VideoTypeToName[vid.Type])
	v.Add("hash", vid.Hash)
	v.Add("time", fmt.Sprintf("%d", vid.Time.Unix()))

	return v.Encode()
}
