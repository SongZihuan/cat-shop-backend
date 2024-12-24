package data

func newSystemError(code CodeType, vargs ...interface{}) Data {
	if code <= 0 {
		panic("code must more than 0")
	}

	if len(vargs) == 0 {
		return newData(code)
	} else if len(vargs) == 1 {
		if err, ok := vargs[0].(error); ok {
			return newData(code, nil, "", err.Error())
		} else if msg, ok := vargs[0].(string); ok {
			return newData(code, nil, "", msg)
		} else {
			panic("bad arguments")
		}
	} else {
		panic("too many arguments")
	}
}

func newClientError(code CodeType, msg string, vargs ...interface{}) Data {
	if code <= 0 {
		panic("code must more than 0")
	}

	if len(vargs) == 0 {
		return newData(code, nil, msg)
	} else if len(vargs) == 1 {
		if err, ok := vargs[0].(error); ok {
			return newData(code, nil, msg, err.Error())
		} else if debugMsg, ok := vargs[0].(string); ok {
			return newData(code, nil, msg, debugMsg)
		} else {
			panic("bad arguments")
		}
	} else {
		panic("too many arguments")
	}
}

func NewSystemUnknownError(dat ...interface{}) Data {
	if len(dat) == 0 {
		return newData(GlobalCodeErrorUnknown)
	} else if len(dat) == 1 {
		if msg, ok := dat[1].(string); ok {
			return newData(GlobalCodeErrorUnknown, nil, "系统错误", msg)
		} else if err, ok := dat[1].(error); ok {
			return newData(GlobalCodeErrorUnknown, nil, "系统错误", err.Error())
		} else {
			return newData(GlobalCodeErrorUnknown)
		}
	} else {
		panic("too many arguments")
	}
}

func NewSystemDataBaseError(err error) Data {
	return newSystemError(GlobalCodeErrorDatabase, err)
}

func NewClientTokenExpireError(debugMsg string) Data {
	return newClientError(GlobalCodeErrorTokenExpire, "登陆过起", debugMsg)
}

func NewClientAdminError() Data {
	return newClientError(GlobalCodeErrorNoPermission, "权限不足")
}

func NewClientRootAdminError() Data {
	return newClientError(GlobalCodeErrorNoRootPermission, "权限不足")
}

func NewClientNotTestError() Data {
	return newClientError(GlobalCodeErrorNotTestMode, "非测试模式")
}

func NewClientBadRequests(err error) Data {
	return newClientError(GlobalCodeErrorNotTestMode, "错误请求", err)
}

func NewClientAdminUserNotFound() Data {
	return newClientError(GlobalCodeErrorAdminApiUserNotFound, "指定操作用户不存在")
}

func NewClientAdminUserNoPermission() Data {
	return newClientError(GlobalCodeErrorAdminApiUserNoPermission, "无权操作指定用户")
}
