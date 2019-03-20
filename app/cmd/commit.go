package root

import (
	"fmt"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes",
	Run: func(cmd *cobra.Command, args []string) {

		commitCommander := run.Commander{
			Command: "git",
			Args:    []string{"commit", "-am", "test commit"},
		}

		result, err := git.CommitChanges(commitCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
