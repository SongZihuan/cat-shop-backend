package model

import "gorm.io/gorm"

type ConfigType string

const (
	String     ConfigType = "string"
	Text       ConfigType = "text"
	Pic        ConfigType = "pic"
	StringMust ConfigType = "string|must"
	TextMust   ConfigType = "text|must"
	PicMust    ConfigType = "pic|must"
)

var ConfigTypeName = map[ConfigType]string{
	StringMust: "必选字符串",
	PicMust:    "必选图片",
	TextMust:   "必选文本域",
	String:     "可选字符串",
	Pic:        "可选图片",
	Text:       "可选文本域",
}

var AdminConfigType = map[string]ConfigType{
	"name":              StringMust,
	"subname":           StringMust,
	"hotline":           StringMust,
	"logo":              PicMust,
	"icon":              PicMust,
	"service":           StringMust,
	"wechat":            PicMust,
	"avatar":            PicMust,
	"aboutus":           StringMust,
	"passwordfronthash": StringMust,
	"footer":            String,
	"adtitle":           StringMust,
	"adpic":             PicMust,
	"ad":                TextMust,
	"adurl":             String,
}

var AdminConfigInfo = map[string]string{
	"name":              "网站标题",
	"subname":           "网站副标题",
	"hotline":           "服务热线",
	"logo":              "网站主Logo",
	"icon":              "网页Icon",
	"service":           "服务标语",
	"wechat":            "微信二维码推广",
	"adtitle":           "广告标题",
	"adpic":             "广告图片",
	"ad":                "广告内容",
	"adurl":             "广告Url",
	"avatar":            "默认头像",
	"aboutus":           "关于我们",
	"passwordfronthash": "密码学加密字段",
	"footer":            "ICP备案号和底部文字内容",
}

type Config struct {
	gorm.Model
	Key   string `gorm:"type:varchar(20);not null"`
	Value string `gorm:"type:varchar(20);not null"`
}
