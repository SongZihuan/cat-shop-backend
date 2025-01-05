package adminupdatexieyi

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

type Query struct {
	Type    modeltype.XieYiType `json:"type"`
	Content string              `json:"content"`
}

type Data struct {
}

func NewData() Data {
	return Data{}
}

func NewJsonData() data.Data {
	return data.NewSuccessWithData(NewData())
}
