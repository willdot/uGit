package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/run"
)

func DeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "d",
		Short: "Delete branches",
		Run: func(cmd *cobra.Command, args []string) {
			err := delete()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}

func delete() error {
	branches, err := git.GetBranches()
	if err != nil {
		return err
	}

	branchesToDelete := askUserToSelectOptions(branches, "Select branches to delete", false)

	if len(branchesToDelete) == 0 {
		fmt.Println("No branches selected")
		return nil
	}

	for _, branch := range branchesToDelete {
		result := deleteBranch(strings.TrimSpace(branch))

		fmt.Println(result)
	}

	return nil
}

func deleteBranch(branch string) string {
	result, err := run.RunCommand("git", []string{"branch", "-d", branch})
	if err != nil {
		handleErrorDelete(result, branch)
		return ""
	}

	return result
}

func handleErrorDelete(errorMessage, branchName string) {
	lines := strings.Split(errorMessage, "\n")
	if len(lines) == 0 {
		return
	}

	if !strings.Contains(lines[1], "If you are sure you want to delete it, run 'git branch -D") {
		fmt.Println(errorMessage)
	}

	res := askUserToSelectSingleOption([]string{"yes", "no"}, "Would you like to force delete this branch?")

	if res == "" {
		return
	}

	if res == "yes" {
		forceDeleteBranch(branchName)
	}
}

func forceDeleteBranch(branch string) {
	result, err := run.RunCommand("git", []string{"branch", "-D", branch})
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
