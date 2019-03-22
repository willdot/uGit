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
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		branchesSlice := git.SplitBranches(branches, true)

		fmt.Println(branchesSlice)

		selection := struct {
			Branch string
		}{}

		question := getQuestion(branchesSlice)
		err = survey.Ask(question, &selection)

		checkout(selection.Branch)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func getQuestion(branches []string) []*survey.Question {
	var selectBranch = []*survey.Question{
		{
			Name: "branch",
			Prompt: &survey.Select{
				Message: "Select a branch",
				Options: branches,
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}

	return selectBranch
}

func checkout(branchSelection string) {

	git.RemoveRemoteOriginFromName(&branchSelection)
	checkoutCommander := run.Commander{
		Command: "git",
		Args:    []string{"checkout", strings.Replace(branchSelection, " ", "", -1)},
	}

	result, err := git.CheckoutBranch(checkoutCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}
