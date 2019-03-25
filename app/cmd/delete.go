package root

import (
	"fmt"
	"os"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
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

		branchesToDelete := askUserToSelectFilesToDelete(branches, "Select the files you wish to delete")

		if len(branchesToDelete) == 0 {
			fmt.Println("No files selected")
			os.Exit(1)
		}

		for _, branch := range branchesToDelete {
			result := deleteBranch(branch)

			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func askUserToSelectFilesToDelete(availableFiles []string, message string) []string {
	options := append([]string{"**Exit and ignore selections**"}, availableFiles...)

	result := []string{}
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}

	survey.AskOne(prompt, &result, nil)

	for i := 0; i < len(result); i++ {
		if result[i] == "**Exit and ignore selections**" {
			return nil
		}
	}

	return result
}

func deleteBranch(branch string) string {

	deleteCommander := run.Commander{
		Command: "git",
		Args:    []string{"branch", "-d", branch},
	}

	status, err := git.DeleteBranch(deleteCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		fmt.Println("end of error")
		return status
	}

	return status
}
