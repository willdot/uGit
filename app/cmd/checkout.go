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

const exit = "Exit"

var checkoutCmd = &cobra.Command{
	Use:   "cko",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {

		branchCommander := run.Commander{
			Command: "git",
			Args:    []string{"branch", "-a"},
		}

		branches, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		branchName := askUser(branches)

		checkout(branchName, false)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func askUser(branches []string) string {

	var branchName string
	question := getBranchQuestion(branches)
	err := survey.Ask(question, &branchName)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	if branchName == exit {
		os.Exit(1)
	}

	return branchName
}

func getBranchQuestion(branches []string) []*survey.Question {

	options := []string{exit}

	options = append(options, branches...)

	var selectBranch = []*survey.Question{
		{
			Name: "branch",
			Prompt: &survey.Select{
				Message: "Select a branch",
				Options: options,
			},
			Validate: survey.Required,
		},
	}

	return selectBranch
}

func checkout(branchSelection string, new bool) {

	git.RemoveRemoteOriginFromName(&branchSelection)

	args := []string{"checkout"}

	if new {
		args = append(args, "-b")
	}

	args = append(args, strings.TrimSpace(branchSelection))

	checkoutCommander := run.Commander{
		Command: "git",
		Args:    args,
	}

	result, err := git.CheckoutBranch(checkoutCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}
