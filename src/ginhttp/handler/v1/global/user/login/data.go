package login

import "github.com/SongZihuan/cat-shop-backend/src/ginhttp/data"

type Query struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Data struct {
	XToken  string `json:"xtoken"`
	Success bool   `json:"success"`
}

func NewData(token ...string) Data {
	if len(token) == 1 && token[0] != "" {
		return Data{
			XToken:  token[0],
			Success: false,
		}
	}
	return Data{
		XToken:  "",
		Success: false,
	}
}

func NewJsonData(token string) data.Data {
	return data.NewSuccessWithData(NewData(token))
}
