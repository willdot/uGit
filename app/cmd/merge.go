package root

import (
	"fmt"
	"strings"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use : "merge",
	Short : "Merge branches",
	Run : func(cmd *cobra.Command, args []string) {
		var branchName string

		branchCommander := run.Commander{
			Command: "git",
			Args:    []string{"branch", "-a"},
		}

		branches, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		branchName = askUser(branches)

		merge(branchName)
	},
}

func merge(branchSelection string) {
	git.RemoveRemoteOriginFromName(&branchSelection)

	args := []string{"merge", strings.TrimSpace(branchSelection)}

	mergeCommander := run.Commander{
		Command: "git",
		Args:    args,
	}

	result, err := git.Merge(mergeCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
