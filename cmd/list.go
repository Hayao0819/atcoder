package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func getContestList() (*[]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	contests := []string{}
	dirs, err := filepath.Glob(path.Join(pwd, "src", "*"))
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		contests = append(contests, path.Base(dir))
	}
	return &contests, nil
}

func list() *cobra.Command {
	cmd := cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			contests, err := getContestList()
			if err != nil {
				return err
			}
			cmd.Println(*contests)
			return nil
		},
	}
	return &cmd
}

func init() {
	cobrautils.AddSubCmds(list())
}
