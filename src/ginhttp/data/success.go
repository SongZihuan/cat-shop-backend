package data

func NewSuccess(msg ...string) Data {
	type data struct {
		Success bool `json:"success"`
	}

	if len(msg) == 0 {
		return newData(GlobalCodeOk, data{Success: true})
	} else if len(msg) == 1 {
		return newData(GlobalCodeOk, data{Success: true}, msg[0])
	} else if len(msg) == 2 {
		return newData(GlobalCodeOk, data{Success: true}, msg[0], msg[1])
	} else {
		panic("too many arguments")
	}
}

func NewSuccessWithData(dat interface{}, msg ...string) Data {
	if len(msg) == 0 {
		return newData(GlobalCodeOk, dat)
	} else if len(msg) == 1 {
		return newData(GlobalCodeOk, dat, msg[0])
	} else if len(msg) == 2 {
		return newData(GlobalCodeOk, dat, msg[0], msg[1])
	} else {
		panic("too many arguments")
	}
}
