package login

import "github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"

const (
	CodePhoneError    data.CodeType = -1
	CodePasswordError data.CodeType = -2
)

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

func NewToken(token string) data.Data {
	return data.NewData(data.GlobalCodeOk, NewData(token))
}

func NewMsgError(code data.CodeType, msg string, debugMsg ...string) data.Data {
	if len(debugMsg) == 0 {
		return data.NewData(code, NewData(), msg)
	} else if len(debugMsg) == 1 {
		return data.NewData(code, NewData(), msg, debugMsg[0])
	} else {
		panic("too many arguments")
	}
}
