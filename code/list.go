package code

import (
	"os"
	"path"

	"github.com/Hayao0819/atcoder/flist"
)

func GetContestList() (*[]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(path.Join(pwd, "src"), flist.WithDirOnly(), flist.WithExactDepth(1), flist.WithFileName())
	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetProblemList(contest string) (*[]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(path.Join(pwd, "src", contest), flist.WithDirOnly(), flist.WithExactDepth(1), flist.WithFileName())
	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetCodeList(contest, problem string) (*[]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(path.Join(pwd, "src", contest, problem), flist.WithFileOnly(), flist.WithExactDepth(1), flist.WithFileName())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetPath(contest, problem, code string) string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "src", contest, problem, code)
}
