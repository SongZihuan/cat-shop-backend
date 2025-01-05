package admingetconfig

import "github.com/SuperH-0630/cat-shop-back/src/model/modeltype"

type ConfigItemType string

const (
	StringMust     ConfigItemType = "string|must"
	PicMust        ConfigItemType = "pic|must"
	TextMust       ConfigItemType = "text|must"
	StringOptional ConfigItemType = "string"
	PicOptional    ConfigItemType = "pic"
	TextOptional   ConfigItemType = "text"
)

// AdminConfigTypeInfo 定义了配置项类型的描述
var AdminConfigTypeInfo = map[ConfigItemType]string{
	StringMust:     "必选字符串",
	PicMust:        "必选图片",
	TextMust:       "必选文本域",
	StringOptional: "可选字符串",
	PicOptional:    "可选图片",
	TextOptional:   "可选文本域",
}

// AdminConfigType 定义了配置项的名称和类型
var AdminConfigType = map[modeltype.ConfigKeyType]ConfigItemType{
	modeltype.KeyName:              StringMust,
	modeltype.KeySubName:           StringMust,
	modeltype.KeyHotline:           StringMust,
	modeltype.KeyLogo:              PicMust,
	modeltype.KeyIcon:              PicMust,
	modeltype.KeyService:           StringMust,
	modeltype.KeyWechat:            PicMust,
	modeltype.KeyAvatar:            PicMust,
	modeltype.KeyAboutUs:           StringMust,
	modeltype.KeyPasswordFrontHash: StringMust,
	modeltype.KeyFooter:            StringOptional,
	modeltype.KeyAdTitle:           StringMust,
	modeltype.KeyAdPic:             PicMust,
	modeltype.KeyAd:                TextMust,
	modeltype.KeyAdUrl:             StringOptional,
}

// AdminConfigInfo 定义了配置项的中文描述
var AdminConfigInfo = map[modeltype.ConfigKeyType]string{
	modeltype.KeyName:              "网站标题",
	modeltype.KeySubName:           "网站副标题",
	modeltype.KeyHotline:           "服务热线",
	modeltype.KeyLogo:              "网站主Logo",
	modeltype.KeyIcon:              "网页Icon",
	modeltype.KeyService:           "服务标语",
	modeltype.KeyWechat:            "微信二维码推广",
	modeltype.KeyAvatar:            "关于我们",
	modeltype.KeyAboutUs:           "密码学加密字段",
	modeltype.KeyPasswordFrontHash: "ICP备案号和底部文字内容",
	modeltype.KeyFooter:            "广告标题",
	modeltype.KeyAdTitle:           "广告图片",
	modeltype.KeyAdPic:             "广告内容",
	modeltype.KeyAd:                "广告Url",
	modeltype.KeyAdUrl:             "默认头像",
}
