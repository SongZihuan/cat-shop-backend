package getconfig

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type Data map[modeltype.ConfigKeyType]modeltype.ConfigValueType

func NewData(cfg []model.Config) Data {
	res := make(Data, len(modeltype.ConfigKey))

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

func NewJsonData(cfg []model.Config) data.Data {
	return data.NewSuccessWithData(NewData(cfg))
}
