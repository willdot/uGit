package root

import (
	"fmt"
	"strings"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

const exit = "Exit"

// checkoutCmd represents the checkout a branch command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
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

		var selection string

		question := getQuestion(branches)
		err = survey.Ask(question, &selection)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		if selection == exit {
			return
		}

		checkout(selection, false)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func getQuestion(branches []string) []*survey.Question {

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
		Args:    args, //[]string{"checkout", strings.TrimSpace(branchSelection)},
	}

	result, err := git.CheckoutBranch(checkoutCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}
