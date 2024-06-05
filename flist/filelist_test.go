package flist_test

import (
	"os"
	"testing"

	"github.com/Hayao0819/atcoder/flist"
)

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestFileList(t *testing.T) {
	home, err := os.UserHomeDir()
	handleError(t, err)

	testcases := []struct {
		options []flist.Option
	}{
		{
			options: []flist.Option{
				flist.WithMaxDepth(1),
				flist.WithRelPath(),
			},
		},
	}

	for _, tc := range testcases {
		list, err := flist.Get(home, tc.options...)
		t.Logf("list: %v", list)
		handleError(t, err)
	}

}
