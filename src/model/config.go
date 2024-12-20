package model

import "gorm.io/gorm"

type ConfigKeyType string
type ConfigValueType string
type TypesOfConfigValueType string

const (
	KeyName              ConfigKeyType = "name"
	KeySubName           ConfigKeyType = "subname"
	KeyHotline           ConfigKeyType = "hotline"
	KeyLogo              ConfigKeyType = "logo"
	KeyIcon              ConfigKeyType = "icon"
	KeyService           ConfigKeyType = "service"
	KeyWechat            ConfigKeyType = "wechat"
	KeyAvatar            ConfigKeyType = "avatar"
	KeyAboutUs           ConfigKeyType = "aboutus"
	KeyPasswordFrontHash ConfigKeyType = "passwordfronthash"
	KeyFooter            ConfigKeyType = "footer"
	KeyAdTitle           ConfigKeyType = "adtitle"
	KeyAdPic             ConfigKeyType = "adpic"
	KeyAd                ConfigKeyType = "ad"
	KeyAdUrl             ConfigKeyType = "adurl"
)

var ConfigKey = []ConfigKeyType{
	KeyName,
	KeySubName,
	KeyHotline,
	KeyLogo,
	KeyIcon,
	KeyService,
	KeyWechat,
	KeyAvatar,
	KeyAboutUs,
	KeyPasswordFrontHash,
	KeyFooter,
	KeyAdTitle,
	KeyAdPic,
	KeyAd,
	KeyAdUrl,
}

var ConfigKeyMap = map[ConfigKeyType]bool{
	KeyName:              true,
	KeySubName:           true,
	KeyHotline:           true,
	KeyLogo:              true,
	KeyIcon:              true,
	KeyService:           true,
	KeyWechat:            true,
	KeyAvatar:            true,
	KeyAboutUs:           true,
	KeyPasswordFrontHash: true,
	KeyFooter:            true,
	KeyAdTitle:           true,
	KeyAdPic:             true,
	KeyAd:                true,
	KeyAdUrl:             true,
}

const (
	TypeString     TypesOfConfigValueType = "string"
	TypeText       TypesOfConfigValueType = "text"
	TypePic        TypesOfConfigValueType = "pic"
	TypeStringMust TypesOfConfigValueType = "string|must"
	TypeTextMust   TypesOfConfigValueType = "text|must"
	TypePicMust    TypesOfConfigValueType = "pic|must"
)

var ConfigTypeName = map[TypesOfConfigValueType]string{
	TypeStringMust: "必选字符串",
	TypePicMust:    "必选图片",
	TypeTextMust:   "必选文本域",
	TypeString:     "可选字符串",
	TypePic:        "可选图片",
	TypeText:       "可选文本域",
}

var ConfigType = map[ConfigKeyType]TypesOfConfigValueType{
	KeyName:              TypeStringMust,
	KeySubName:           TypeStringMust,
	KeyHotline:           TypeStringMust,
	KeyLogo:              TypePicMust,
	KeyIcon:              TypePicMust,
	KeyService:           TypeStringMust,
	KeyWechat:            TypePicMust,
	KeyAvatar:            TypePicMust,
	KeyAboutUs:           TypeStringMust,
	KeyPasswordFrontHash: TypeStringMust,
	KeyFooter:            TypeString,
	KeyAdTitle:           TypeStringMust,
	KeyAdPic:             TypePicMust,
	KeyAd:                TypeTextMust,
	KeyAdUrl:             TypeString,
}

var ConfigInfo = map[ConfigKeyType]string{
	KeyName:              "网站标题",
	KeySubName:           "网站副标题",
	KeyHotline:           "服务热线",
	KeyLogo:              "网站主Logo",
	KeyIcon:              "网页Icon",
	KeyService:           "服务标语",
	KeyWechat:            "微信二维码推广",
	KeyAvatar:            "默认头像",
	KeyAboutUs:           "关于我们",
	KeyPasswordFrontHash: "密码学加密字段",
	KeyFooter:            "ICP备案号和底部文字内容",
	KeyAdTitle:           "广告标题",
	KeyAdPic:             "广告图片",
	KeyAd:                "广告内容",
	KeyAdUrl:             "广告Url",
}

type Config struct {
	gorm.Model
	Key   ConfigKeyType          `gorm:"type:varchar(20);not null;uniqueIndex"`
	Value ConfigValueType        `gorm:"type:varchar(20);not null"`
	Type  TypesOfConfigValueType `gorm:"type:varchar(20);not null"`
}
