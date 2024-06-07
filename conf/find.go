package conf

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Hayao0819/nahi/futils"
)

var configName = ".atcoderrc.json"

func FindBaseDir() (*string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	targetDir := wd
	for {
		if hasConfigFile(targetDir) {
			return &targetDir, nil
		}
		targetDir = filepath.Dir(targetDir)
		if targetDir == "/" {
			return nil, errors.New("you are not in the atcoder project directory")
		}
	}
}

func hasConfigFile(dir string) bool {
	return futils.Exists(filepath.Join(dir, configName))
}
