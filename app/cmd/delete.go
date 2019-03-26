package root

import (
	"fmt"
	"os"
	"strings"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "d",
	Short: "Delete branches",
	Run: func(cmd *cobra.Command, args []string) {

		branchCommander := run.Commander{
			Command: "git",
			Args:    []string{"branch"},
		}

		branches, err := git.GetBranches(branchCommander)

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

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteBranch(branch string) string {

	deleteCommander := run.Commander{
		Command: "git",
		Args:    []string{"branch", "-d", branch},
	}

	result, err := git.DeleteBranch(deleteCommander)

	if err != nil {
		handleErrorDelete(result, branch)
		return ""
	}

	return result
}

func forceDeleteBranch(branch string) {
	deleteCommander := run.Commander{
		Command: "git",
		Args:    []string{"branch", "-D", branch},
	}

	result, _ := git.DeleteBranch(deleteCommander)
	fmt.Println(result)
}

func handleErrorDelete(errorMessage, branchName string) {

	lines := strings.Split(errorMessage, "\n")

	if strings.HasPrefix(lines[1], "If you are sure you want to delete it, run 'git branch -D") {

		fmt.Println(lines[0])
		result := false

		prompt := &survey.Confirm{
			Message: "Would you like to force delete this branch?",
		}

		survey.AskOne(prompt, &result, nil)

		if result == true {
			forceDeleteBranch(branchName)
		}
	} else {
		fmt.Println(errorMessage)
	}
}
