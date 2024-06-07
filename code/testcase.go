package code

import "github.com/Hayao0819/atcoder/flist"

func GetTestCases(contest, problem string) (*[]string, error) {
	problemDir := GetPath(contest, problem, "")
	files, err := flist.Get(problemDir, flist.WithFileOnly(), flist.WithExtOnly(".txt"))
	if err != nil {
		return nil, err
	}
	return files, nil
}
