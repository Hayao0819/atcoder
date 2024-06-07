package conf

import (
	"errors"
	"path"
)

var contestDir = ""

func GetContestDir() (string, error) {
	if contestDir == "" {
		return "", errors.New("not initialized")
	}
	return contestDir, nil
}

func ReadConfig() error {
	return nil
}

func init() {
	baseDir, err := FindBaseDir()
	if err != nil {
		return
	}
	contestDir = path.Join(*baseDir, "src")
}
