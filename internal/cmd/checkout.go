package root

import (
	"fmt"
	"uGit/internal/pkg/git"
	"uGit/internal/pkg/run"

	"github.com/spf13/cobra"
)

// checkoutCmd represents the say command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select a branch")

		commander := run.RealCommander{}
		result, _ := run.CommandWithResult(commander, "git", "branch")
		fmt.Println(result)

		branches := git.SplitBranches(result)

		for _, branch := range branches {
			fmt.Println(branch)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
