package data

type CodeType int64

const (
	GlobalCodeOk                    CodeType = 0
	GlobalCodeErrorUnknown          CodeType = 1
	GlobalCodeErrorTokenExpire      CodeType = 2
	GlobalCodeErrorNoPermission     CodeType = 3
	GlobalCodeErrorNoRootPermission CodeType = 4
	GlobalCodeErrorNotTestMode      CodeType = 5
	GlobalCodeErrorDatabase         CodeType = 6
	GlobalCodeErrorBadRequests      CodeType = 7
)

var CodeDebugMsg = map[CodeType]string{
	GlobalCodeOk:                    "正常",
	GlobalCodeErrorUnknown:          "未知错误",
	GlobalCodeErrorTokenExpire:      "Token过期",
	GlobalCodeErrorNoPermission:     "无访问权限（Admin）",
	GlobalCodeErrorNoRootPermission: "无访问权限（Root Admin）",
	GlobalCodeErrorNotTestMode:      "非测试模式",
	GlobalCodeErrorDatabase:         "数据库错误",
	GlobalCodeErrorBadRequests:      "错误请求",
}

func GetCodeDebugMsg(code CodeType) string {
	if code < 0 {
		return "业务错误"
	}

	debugMsg, ok := CodeDebugMsg[code]
	if ok {
		return debugMsg
	}
	return "未知错误"
}
