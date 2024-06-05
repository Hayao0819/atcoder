package code

import (
	"errors"
	"os"

	"github.com/Hayao0819/nahi/exutils"
)

func isExecutable(mode os.FileMode) bool {
	return mode&0111 != 0
}
func Execute(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !isExecutable(info.Mode()) {
		return errors.New("file is not executable")
	}

	cmd := exutils.CommandWithStdio(path)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
