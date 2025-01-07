package adminaddwupin

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
	"mime/multipart"
)

type Query struct {
	Name      string                  `json:"name"`
	ClassID   uint                    `json:"classId"`
	Tag       string                  `json:"tag"`
	HotPrice  modeltype.PriceNullJson `json:"hotPrice"`
	RealPrice modeltype.Price         `json:"realPrice"`
	Info      string                  `json:"info"`
	Ren       string                  `json:"ren"`
	Phone     string                  `json:"phone"`
	Email     string                  `json:"email"`
	Wechat    string                  `json:"wechat"`
	Location  string                  `json:"location"`
	Pic       string                  `json:"pic"`
	Hot       bool                    `json:"hot"`
	Down      bool                    `json:"down"`
	File      *multipart.FileHeader   `form:"file"`
}
