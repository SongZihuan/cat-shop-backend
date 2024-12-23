package data

type CodeType int64

func GetCodeName(code CodeType) string {
	if code < 0 {
		return "业务错误"
	} else if code == 0 {
		return "正常"
	} else {
		debugMsg, ok := GlobalErrorCodeName[code]
		if ok {
			return debugMsg
		}

		return "未知错误"
	}
}
