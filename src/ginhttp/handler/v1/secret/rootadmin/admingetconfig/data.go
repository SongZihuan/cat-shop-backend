package admingetconfig

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

type Config map[modeltype.ConfigKeyType]modeltype.ConfigValueType

func NewConfig(cfg []model.Config) Config {
	res := make(Config, len(modeltype.ConfigKey))

	for _, k := range modeltype.ConfigKey {
		res[k] = ""
	}

	for _, r := range cfg {
		_, ok := res[r.Key] // 检测是否存在，用于确定r.Key是否存在于ConfigKey
		if ok {
			res[r.Key] = r.GetValue()
		}
	}

	return res
}

type Data struct {
	Config   Config                                                       `json:"config"`
	Info     map[modeltype.ConfigKeyType]string                           `json:"info"`
	Type     map[modeltype.ConfigKeyType]modeltype.TypesOfConfigValueType `json:"type"`
	TypeName map[modeltype.TypesOfConfigValueType]string                  `json:"typename"`
}

func NewData(cfg []model.Config) Data {
	return Data{
		Config:   NewConfig(cfg),
		Info:     modeltype.ConfigInfo,
		Type:     modeltype.ConfigType,
		TypeName: modeltype.ConfigTypeName,
	}
}

func NewJsonData(cfg []model.Config) data.Data {
	return data.NewSuccessWithData(NewConfig(cfg))
}
