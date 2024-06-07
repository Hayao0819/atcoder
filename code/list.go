package code

import (
	"path"

	"github.com/Hayao0819/atcoder/conf"
	"github.com/Hayao0819/atcoder/flist"
)

func GetContestList() (*[]string, error) {
	contestDir, err := conf.GetContestDir()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(contestDir, flist.WithDirOnly(), flist.WithExactDepth(1), flist.WithFileName())
	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetProblemList(contest string) (*[]string, error) {
	contestDir, err := conf.GetContestDir()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(path.Join(contestDir, contest), flist.WithDirOnly(), flist.WithExactDepth(1), flist.WithFileName())
	if err != nil {
		return nil, err
	}

	return files, nil
}

func GetCodeList(contest, problem string) (*[]string, error) {
	contestDir, err := conf.GetContestDir()
	if err != nil {
		return nil, err
	}

	files, err := flist.Get(path.Join(contestDir, contest, problem), flist.WithFileOnly(), flist.WithExactDepth(1), flist.WithFileName(), flist.WithExecutableOnly())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetPath(contest, problem, code string) (string, error) {
	contestDir, err := conf.GetContestDir()
	if err != nil {
		return "", err
	}
	return path.Join(contestDir, contest, problem, code), nil
}
