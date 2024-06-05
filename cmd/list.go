package cmd

import (
	"github.com/Hayao0819/atcoder/code"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func list() *cobra.Command {
	cmd := cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			contests, err := code.GetContestList()
			if err != nil {
				return err
			}
			for _, c := range *contests {
				cmd.Println(c)
			}
			return nil
		},
	}
	return &cmd
}

func init() {
	cobrautils.RegisterSubCmd(list())
}
