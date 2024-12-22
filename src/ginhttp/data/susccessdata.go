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

func NewSuccessData(msg string, debugMsg string) Data {
	return NewData(GlobalCodeOk, NewSuccess(true), msg, debugMsg)
}

func NewNotSuccessData(code CodeType, msg string, debugMsg string) Data {
	return NewData(code, NewSuccess(false), msg, debugMsg)
}
