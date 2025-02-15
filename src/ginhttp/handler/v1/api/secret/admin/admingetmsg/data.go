package admingetmsg

import (
	"github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"
	"github.com/SongZihuan/cat-shop-backend/src/model"
)

type Query struct {
	Page     int `form:"page"`
	PageSize int `form:"pagesize"`
}

type Msg struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"userId"`
	Msg    string `json:"msg"`
	Time   int64  `json:"time"`
}

func NewMsg(msg model.Msg) Msg {
	return Msg{
		ID:     msg.ID,
		UserID: msg.UserID,
		Msg:    msg.Msg,
		Time:   msg.Time.Unix(),
	}
}

type Data struct {
	List     []Msg `json:"list"`
	Total    int   `json:"total"`
	MaxCount int   `json:"maxpage"`
}

func NewData(m []model.Msg, maxcount int) Data {
	res := make([]Msg, 0, len(m))
	for _, v := range m {
		res = append(res, NewMsg(v))
	}

	return Data{
		List:     res,
		Total:    len(res),
		MaxCount: maxcount,
	}
}

func NewJsonData(m []model.Msg, maxcount int) data.Data {
	return data.NewSuccessWithData(NewData(m, maxcount))
}
