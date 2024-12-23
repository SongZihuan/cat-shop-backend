package testpay

import "github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"

type Query struct {
	ID uint `form:"id"`
}

type Data struct {
}

func NewData() Data {
	return Data{}
}

func NewJsonData() data.Data {
	return data.NewSuccessWithData(NewData())
}
