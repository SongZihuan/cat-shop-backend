package admingetconfig

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type ConfigKeyValue map[modeltype.ConfigKeyType]modeltype.ConfigValueType

type Data struct {
	Config   ConfigKeyValue                             `json:"config"`
	Info     map[modeltype.ConfigKeyType]string         `json:"info"`
	Type     map[modeltype.ConfigKeyType]ConfigItemType `json:"type"`
	TypeName map[ConfigItemType]string                  `json:"typename"`
}

func NewConfigKeyValue(cfg []model.Config) ConfigKeyValue {
	config := make(ConfigKeyValue, len(modeltype.ConfigKey))
	for _, k := range modeltype.ConfigKey {
		config[k] = ""
	}

	for _, r := range cfg {
		_, ok := config[r.Key] // 检测是否存在，用于确定r.Key是否存在于ConfigKey
		if ok {
			config[r.Key] = r.GetValue()
		}
	}

	return config
}

func NewData(cfg []model.Config) Data {
	return Data{
		Config:   NewConfigKeyValue(cfg),
		Info:     AdminConfigInfo,
		Type:     AdminConfigType,
		TypeName: AdminConfigTypeInfo,
	}
}

func NewJsonData(cfg []model.Config) data.Data {
	return data.NewSuccessWithData(NewData(cfg))
}
