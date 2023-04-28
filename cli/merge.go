package cli

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/run"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge branches",
	Run: func(cmd *cobra.Command, args []string) {
		var branchName string

		// branchCommander := run.Commander{
		// 	Command: "git",
		// 	Args:    []string{"branch", "-a"},
		// }

		branches, err := git.GetBranches()

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		branchName = askUserToSelectSingleOption(branches, "")
		if branchName == "" {
			return
		}

		merge(branchName)
	},
}

func merge(branchSelection string) {
	git.RemoveRemoteOriginFromName(&branchSelection)

	args := []string{"merge", strings.TrimSpace(branchSelection)}

	// mergeCommander := run.Commander{
	// 	Command: "git",
	// 	Args:    args,
	// }

	// result, err := git.Merge(mergeCommander)
	result, err := run.RunCommand("git", args)
	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}

func something() {

}
