package data

import "github.com/SongZihuan/cat-shop-backend/src/config"

type Data struct {
	Code     CodeType    `json:"code"`
	Data     interface{} `json:"data"`
	Msg      string      `json:"msg"`
	DebugMsg string      `json:"debugmsg,omitempty"`
}

func _newData(code CodeType, vargs ...interface{}) Data {
	if len(vargs) == 0 {
		return Data{Code: code, Data: nil, Msg: GetCodeName(code), DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 1 {
		return Data{Code: code, Data: vargs[0], Msg: GetCodeName(code), DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 2 {
		return Data{Code: code, Data: vargs[0], Msg: vargs[1].(string), DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 3 {
		return Data{Code: code, Data: vargs[0], Msg: vargs[1].(string), DebugMsg: vargs[2].(string)}
	} else {
		panic("too many arguments")
	}
}

func newData(code CodeType, vargs ...interface{}) Data {
	if !config.IsReady() {
		panic("config is not ready")
	}

	dat := _newData(code, vargs...)
	if config.Config().Yaml.Http.DebugMsg.Is(config.Disable) {
		dat.DebugMsg = ""
	}

	return dat
}
