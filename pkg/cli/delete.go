package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/run"
)

func DeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "d",
		Short: "Delete branches",
		Run: func(cmd *cobra.Command, args []string) {
			branches, err := git.GetBranches()
			if err != nil {
				fmt.Printf("error: %v", errors.WithMessage(err, ""))
				return
			}

			branchesToDelete := askUserToSelectOptions(branches, "Select branches to delete", false)

			if len(branchesToDelete) == 0 {
				fmt.Println("No branches selected")
				os.Exit(1)
			}

			for _, branch := range branchesToDelete {
				result := deleteBranch(strings.TrimSpace(branch))

				fmt.Println(result)
			}
		},
	}

	return cmd
}

func deleteBranch(branch string) string {
	result, err := run.RunCommand("git", []string{"branch", "-d", branch})
	if err != nil {
		handleErrorDelete(result, branch)
		return ""
	}

	return result
}

func forceDeleteBranch(branch string) {
	result, err := run.RunCommand("git", []string{"branch", "-D", branch})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func handleErrorDelete(errorMessage, branchName string) {

	lines := strings.Split(errorMessage, "\n")

	if strings.HasPrefix(lines[1], "If you are sure you want to delete it, run 'git branch -D") {

		fmt.Println(lines[0])
		result := false

		res := askUserToSelectSingleOption([]string{"yes", "no"}, "Would you like to force delete this branch?")

		if res == "" {
			return
		}

		if res == "yes" {
			result = true
		}

		if result == true {
			forceDeleteBranch(branchName)
		}
	} else {
		fmt.Println(errorMessage)
	}
}
