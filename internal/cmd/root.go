package root

import (
	"fmt"
	"os"
	"uGit/internal/pkg/git"
	"uGit/internal/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uGit",
	Short: "uGit is an application to run common git commands quicker ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting")

		commander := run.RealCommander{}
		result, _ := run.CommandWithResult(commander, "git", "branch")

		branches := git.SplitBranches(result)
		current, err := git.GetCurrentBranch(branches)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
		}

		fmt.Println(current)
	},
}

// Execute starts the app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
