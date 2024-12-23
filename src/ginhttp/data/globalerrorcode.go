package data

const (
	GlobalCodeErrorUnknown          CodeType = 1
	GlobalCodeErrorTokenExpire      CodeType = 2
	GlobalCodeErrorNoPermission     CodeType = 3
	GlobalCodeErrorNoRootPermission CodeType = 4
	GlobalCodeErrorNotTestMode      CodeType = 5
	GlobalCodeErrorDatabase         CodeType = 6
	GlobalCodeErrorBadRequests      CodeType = 7
)

var GlobalErrorCodeName = map[CodeType]string{ // 不包含GlobalCodeOK
	GlobalCodeErrorUnknown:          "未知错误",
	GlobalCodeErrorTokenExpire:      "Token过期",
	GlobalCodeErrorNoPermission:     "无访问权限（Admin）",
	GlobalCodeErrorNoRootPermission: "无访问权限（Root Admin）",
	GlobalCodeErrorNotTestMode:      "非测试模式",
	GlobalCodeErrorDatabase:         "数据库错误",
	GlobalCodeErrorBadRequests:      "错误请求",
}
