package error

import "errors"

type BuyRecordStatusError struct {
	msg      string
	debugMsg string
}

func NewBuyRecordStatusError(msg string, debugMsg ...string) *BuyRecordStatusError {
	if len(debugMsg) == 0 {
		return &BuyRecordStatusError{msg: msg}
	} else if len(debugMsg) == 1 {
		return &BuyRecordStatusError{msg: msg, debugMsg: debugMsg[0]}
	} else {
		panic("too many arguments")
	}
}

func IsBuyRecordStatusError(err error) (*BuyRecordStatusError, bool) {
	var statusErr *BuyRecordStatusError
	ok := errors.As(err, &statusErr)
	if ok {
		return statusErr, true
	}
	return nil, false
}

func (err *BuyRecordStatusError) Error() string {
	return err.msg
}

func (err *BuyRecordStatusError) DebugMsg() string {
	return err.debugMsg
}
