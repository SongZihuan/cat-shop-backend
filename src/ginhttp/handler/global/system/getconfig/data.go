package getconfig

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
)

type Data map[model.ConfigKeyType]model.ConfigValueType

func NewData(cfg []model.Config) Data {
	res := make(Data, len(model.ConfigKey))

	for _, k := range model.ConfigKey {
		res[k] = ""
	}

	for _, r := range cfg {
		_, ok := res[r.Key] // 检测是否存在，用于确定r.Key是否存在于ConfigKey
		if ok {
			res[r.Key] = r.Value
		}
	}

	return res
}

func NewJsonData(cfg []model.Config) data.Data {
	return data.NewData(data.GlobalCodeOk, NewData(cfg))
}
