package admingetclass

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
)

type Query struct {
	Page     int `form:"page"`
	PageSize int `form:"pagesize"`
}

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Show bool   `json:"show"`
	Down bool   `json:"down"`
}

type Data struct {
	List     []Class `json:"list"`
	Total    int     `json:"total"`
	MaxCount int     `json:"maxpage"`
}

func NewData(list []model.Class, maxcount int) Data {
	res := make([]Class, 0, len(list))
	for _, v := range list {
		if v.Show {
			if v.ID == modeltype.ClassEmptyID {
				res = append(res, Class{
					ID:   modeltype.ClassEmptyID,
					Name: modeltype.ClassEmptyName,
					Show: modeltype.ClassEmptyShow,
					Down: modeltype.ClassEmptyDown,
				})
			} else {
				res = append(res, Class{
					ID:   v.ID,
					Name: v.Name,
					Show: v.Show,
					Down: v.Down,
				})
			}
		}
	}

	return Data{
		List:     res,
		Total:    len(list),
		MaxCount: maxcount,
	}
}

func NewJsonData(list []model.Class, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(list, maxcount))
}
