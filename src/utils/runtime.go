package utils

import (
	"path/filepath"
	"runtime"
)

func GetCallingFunctionInfo() (funcName, file string, fileName string, line int) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "", "", "", 0
	}

	// 获取函数名
	funcName = runtime.FuncForPC(pc).Name()

	// 去除函数地址前缀，得到纯净的函数名
	funcName = runtime.FuncForPC(pc).Name()

	// 获取文件名和行号
	_, file, line, _ = runtime.Caller(1)

	return funcName, file, filepath.Base(file), line
}
