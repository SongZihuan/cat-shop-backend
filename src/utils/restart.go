package utils

import (
	"os"
	"os/exec"
)

func Restart(newArgs ...string) error {
	args := make([]string, 0, len(os.Args))
	copy(args, os.Args)

	if len(newArgs) != 0 {
		args = append(args, newArgs...)
	}

	args0, err := os.Executable()
	if err != nil {
		return err
	}

	err = exec.Command(args0, args[1:]...).Run()
	if err != nil {
		return err
	}

	return nil
}
