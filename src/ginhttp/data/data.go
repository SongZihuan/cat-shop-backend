package data

type Data struct {
	Code     CodeType    `json:"code"`
	Data     interface{} `json:"data"`
	Msg      string      `json:"msg"`
	DebugMsg string      `json:"debugmsg"`
}

func newData(code CodeType, vargs ...interface{}) Data {
	if len(vargs) == 0 {
		return Data{Code: code, Data: nil, Msg: "", DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 1 {
		return Data{Code: code, Data: vargs[0], Msg: "", DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 2 {
		return Data{Code: code, Data: vargs[0], Msg: vargs[1].(string), DebugMsg: GetCodeName(code)}
	} else if len(vargs) == 3 {
		return Data{Code: code, Data: vargs[0], Msg: vargs[1].(string), DebugMsg: vargs[2].(string)}
	} else {
		panic("too many arguments")
	}
}
