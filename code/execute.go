package code

import (
	"errors"
	"os"
	"os/exec"
)

func isExecutable(mode os.FileMode) bool {
	return mode&0111 != 0
}
func Execute(exe, stdin string) error {
	info, err := os.Stat(exe)
	if err != nil {
		return err
	}
	if !isExecutable(info.Mode()) {
		return errors.New("file is not executable")
	}

	//cmd := exutils.CommandWithStdio(exe)

	cmd := exec.Command(exe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	file, err := os.Open(stdin)
	if err != nil {
		return nil
	}
	cmd.Stdin = file
	defer file.Close()

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
