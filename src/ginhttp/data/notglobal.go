package data

type Success struct {
	Success bool `json:"success"`
}

func NewSuccess(success ...bool) Success {
	if len(success) == 0 {
		return Success{
			Success: true,
		}
	} else if len(success) == 1 {
		return Success{
			Success: success[0],
		}
	} else {
		panic("too many arguments")
	}
}

func NewSuccessData(msg string, debugMsg ...string) Data {
	if len(debugMsg) == 0 {
		return NewData(GlobalCodeOk, NewSuccess(true), msg)
	} else if len(debugMsg) == 1 {
		return NewData(GlobalCodeOk, NewSuccess(true), msg, debugMsg[0])
	} else {
		panic("too many arguments")
	}
}

func NewNotSuccessData(code CodeType, msg string, debugMsg ...string) Data {
	if code >= 0 {
		panic("code must less than 0")
	}

	if len(debugMsg) == 0 {
		return NewData(code, NewSuccess(false), msg)
	} else if len(debugMsg) == 1 {
		return NewData(code, NewSuccess(false), msg, debugMsg[0])
	} else {
		panic("too many arguments")
	}
}
