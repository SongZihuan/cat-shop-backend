package mainfunc

import (
	"fmt"
	"os"
)

func exitByError(err error) {
	if err == nil {
		exitByMsg("")
		return
	}

	exitByMsg(err.Error())
}

func exitByMsg(msg string) {
	if len(msg) == 0 {
		msg = "exit: unknown error"
	}

	fmt.Println(msg)
	exit(1)
}

func exit(code ...int) {
	fmt.Println("error exit...")
	if len(code) == 1 {
		os.Exit(code[0])
	}
	os.Exit(0)
}
