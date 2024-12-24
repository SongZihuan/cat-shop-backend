package httpstop

var stopChan = make(chan bool)

func GetStopChan() chan bool {
	return stopChan
}

func SetStop() {
	stopChan <- true
}
