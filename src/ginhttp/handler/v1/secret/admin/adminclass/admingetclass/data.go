package admingetclass

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
)

type Query struct {
	ID uint `json:"ID"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Show bool   `json:"show"`
	Down bool   `json:"down"`
}

func NewClassEmptyData() Class {
	return Class{
		ID:   modeltype.ClassEmptyID,
		Name: modeltype.ClassEmptyName,
		Show: modeltype.ClassEmptyShow,
		Down: modeltype.ClassEmptyDown,
	}
}

func NewData(cls *model.Class) Class {
	if cls.ID == modeltype.ClassEmptyID {
		return NewClassEmptyData()
	}

	return Class{
		ID:   cls.ID,
		Name: cls.Name,
		Show: cls.Show,
		Down: cls.ClassDown,
	}
}

func NewJsonData(cls *model.Class) data.Data {
	return data.NewSuccessWithData(NewData(cls))
}

func NewClassEmptyJsonData() data.Data {
	return data.NewSuccessWithData(NewClassEmptyData())
}
