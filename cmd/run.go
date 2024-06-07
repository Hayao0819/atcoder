package cmd

import (
	"fmt"
	"path"

	"log"

	"github.com/Hayao0819/atcoder/code"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "run",
		Args: cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			contest := ""
			problem := ""
			file := ""

			if len(args) > 0 {
				contest = args[0]
			}
			if len(args) > 1 {
				problem = args[1]
			}
			if len(args) > 2 {
				file = args[2]
			}

			if contest == "" {
				contents, err := code.GetContestList()
				if err != nil {
					return err
				}
				prompt := promptui.Select{
					Label: "Select contest",
					Items: *contents,
				}

				_, result, err := prompt.Run()
				if err != nil {
					return err
				}
				contest = result
			}

			if problem == "" {
				problems, err := code.GetProblemList(contest)
				if err != nil {
					return err
				}

				if len(*problems) == 0 {
					return fmt.Errorf("no problem found")
				}

				if len(*problems) == 1 {
					problem = (*problems)[0]
				} else {

					prompt := promptui.Select{
						Label: "Select problem",
						Items: *problems,
					}

					_, result, err := prompt.Run()
					if err != nil {
						return err
					}
					problem = result
				}
			}

			if file == "" {
				codes, err := code.GetCodeList(contest, problem)
				if err != nil {
					return err
				}
				if len(*codes) == 0 {
					return fmt.Errorf("no code found")
				}

				if len(*codes) == 1 {
					file = (*codes)[0]
				} else {
					prompt := promptui.Select{
						Label: "Select code",
						Items: *codes,
					}

					_, result, err := prompt.Run()
					if err != nil {
						return err
					}
					file = result
				}
			}

			// Get target script path
			fullpath := code.GetPath(contest, problem, file)

			// Get test cases
			testcases, err := code.GetTestCases(contest, problem)
			if err != nil {
				return err
			}

			// Execute target script with test cases
			if len(*testcases) == 0 {
				return fmt.Errorf("no test case found")
			}
			for _, testcase := range *testcases {
				log.Printf("Execute %s-%s with %s", contest, problem, path.Base(testcase))
				if err := code.Execute(fullpath, testcase); err != nil {
					return err
				}
			}
			return nil
		},
	}
	return &cmd
}

func init() {
	cobrautils.RegisterSubCmd(runCmd())
}
