package data

func NewSystemError(code CodeType, vargs ...interface{}) Data {
	if len(vargs) == 0 {
		return NewData(code)
	} else if len(vargs) == 1 {
		if err, ok := vargs[0].(error); ok {
			return NewData(code, nil, "", err.Error())
		} else if msg, ok := vargs[0].(string); ok {
			return NewData(code, nil, "", msg)
		} else {
			panic("bad arguments")
		}
	} else {
		panic("too many arguments")
	}
}

func NewSystemUnknownError(dat ...interface{}) Data {
	if len(dat) == 0 {
		return NewData(GlobalCodeErrorUnknown)
	} else if len(dat) == 1 {
		if msg, ok := dat[1].(string); ok {
			return NewData(GlobalCodeErrorUnknown, nil, "系统错误", msg)
		} else if err, ok := dat[1].(error); ok {
			return NewData(GlobalCodeErrorUnknown, nil, "系统错误", err.Error())
		} else {
			return NewData(GlobalCodeErrorUnknown)
		}
	} else {
		panic("too many arguments")
	}
}

func NewSystemDataBaseError(err error) Data {
	return NewSystemError(GlobalCodeErrorDatabase, err)
}

func NewClientError(code CodeType, msg string, vargs ...interface{}) Data {
	if len(vargs) == 0 {
		return NewData(code, nil, msg)
	} else if len(vargs) == 1 {
		if err, ok := vargs[0].(error); ok {
			return NewData(code, nil, msg, err.Error())
		} else if debugMsg, ok := vargs[0].(string); ok {
			return NewData(code, nil, msg, debugMsg)
		} else {
			panic("bad arguments")
		}
	} else {
		panic("too many arguments")
	}
}

func NewClientTokenExpireError(debugMsg string) Data {
	return NewClientError(GlobalCodeErrorTokenExpire, "登陆过起", debugMsg)
}

func NewClientAdminError() Data {
	return NewClientError(GlobalCodeErrorNoPermission, "权限不足")
}

func NewClientRootAdminError() Data {
	return NewClientError(GlobalCodeErrorNoRootPermission, "权限不足")
}

func NewClientNotTestError() Data {
	return NewClientError(GlobalCodeErrorNotTestMode, "非测试模式")
}
