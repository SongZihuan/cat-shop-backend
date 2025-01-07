package getclasslst

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
)

type Class struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	List  []Class `json:"list"`
	Total int     `json:"total"`
}

func NewData(list []model.Class) Data {
	res := make([]Class, 0, len(list))
	for _, v := range list {
		if v.IsClassShow() {
			res = append(res, Class{
				ID:   v.ID,
				Name: v.Name,
			})
		}
	}

	return Data{
		List:  res,
		Total: len(list),
	}
}

func NewJsonData(list []model.Class) data.Data {
	return data.NewSuccessWithData(NewData(list))
}
