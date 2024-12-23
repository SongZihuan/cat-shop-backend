package data

func NewCustomError(code CodeType, msg string, debugMsg ...string) Data {
	if code >= 0 {
		panic("code must less than 0")
	}

	type data struct {
		Success bool `json:"success"`
	}

	if len(debugMsg) == 0 {
		return newData(code, data{Success: false}, msg)
	} else if len(debugMsg) == 1 {
		return newData(code, data{Success: false}, msg, debugMsg[0])
	} else {
		panic("too many arguments")
	}
}
