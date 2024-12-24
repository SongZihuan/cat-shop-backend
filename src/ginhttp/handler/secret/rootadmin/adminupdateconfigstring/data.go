package adminupdateconfigstring

import "github.com/SuperH-0630/cat-shop-back/src/model/modeltype"

type Query struct {
	Key   modeltype.ConfigKeyType   `form:"key"`
	Value modeltype.ConfigValueType `form:"value"`
}
