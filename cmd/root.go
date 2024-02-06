package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use: "atcoder-tools",
	}
	cobrautils.AddSubCmdsToRoot(&cmd)
	return &cmd
}
