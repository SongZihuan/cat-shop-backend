package data

const (
	GlobalCodeErrorUnknown                  CodeType = 1
	GlobalCodeErrorTokenExpire              CodeType = 2
	GlobalCodeErrorNoPermission             CodeType = 3
	GlobalCodeErrorNoRootPermission         CodeType = 4
	GlobalCodeErrorNotTestMode              CodeType = 5
	GlobalCodeErrorDatabase                 CodeType = 6
	GlobalCodeErrorBadRequests              CodeType = 7
	GlobalCodeErrorAdminApiUserNotFound     CodeType = 8
	GlobalCodeErrorAdminApiUserNoPermission CodeType = 9
	GlobalCodeErrNotFound                   CodeType = 10
	GlobalCodeErrNotCors                    CodeType = 11
)

var GlobalErrorCodeName = map[CodeType]string{ // 不包含GlobalCodeOK
	GlobalCodeErrorUnknown:                  "未知错误",
	GlobalCodeErrorTokenExpire:              "Token过期",
	GlobalCodeErrorNoPermission:             "无访问权限API（Admin）",
	GlobalCodeErrorNoRootPermission:         "无访问权限API（Root Admin）",
	GlobalCodeErrorNotTestMode:              "非测试模式",
	GlobalCodeErrorDatabase:                 "数据库错误",
	GlobalCodeErrorBadRequests:              "错误请求",
	GlobalCodeErrorAdminApiUserNotFound:     "指定操作用户不存在",
	GlobalCodeErrorAdminApiUserNoPermission: "无权操作指定用户",
	GlobalCodeErrNotFound:                   "资源不存在",
	GlobalCodeErrNotCors:                    "不允许跨域",
}
