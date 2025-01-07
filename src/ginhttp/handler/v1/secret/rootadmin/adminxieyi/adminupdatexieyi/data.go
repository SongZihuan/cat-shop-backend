package adminupdatexieyi

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
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
