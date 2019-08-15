package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/app/pkg/git"
	"github.com/willdot/uGit/app/pkg/run"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

const exit = "Exit"

var newBranchFlag bool

var checkoutCmd = &cobra.Command{
	Use:   "cko",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {

		var branchName string

		if newBranchFlag {
			question := getBranchNameQuestion()
			err := survey.Ask(question, &branchName)

			if err != nil {
				fmt.Printf("error: %v", errors.WithMessage(err, ""))
				return
			}

			if branchName == "exit" {
				os.Exit(1)
			}

		} else {

			branchCommander := run.Commander{
				Command: "git",
				Args:    []string{"branch", "-a"},
			}

			branches, err := git.GetBranches(branchCommander)

			if err != nil {
				fmt.Printf("error: %v", errors.WithMessage(err, ""))
				return
			}

			branchName = askUser(branches)
		}

		checkout(branchName, newBranchFlag)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolVarP(&newBranchFlag, "new", "n", false, "create new branch")
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

func getBranchNameQuestion() []*survey.Question {
	var branchNameMessage = []*survey.Question{
		{
			Name: "branch name",
			Prompt: &survey.Input{
				Message: `Enter a branch name or "exit" to cancel`,
			},
			Validate: func(val interface{}) error {

				input := strings.Replace(val.(string), " ", "", -1)

				if str, ok := val.(string); !ok || len(str) != len(input) {
					return errors.New("Branch name cannot contain spaces")
				}
				if str, ok := val.(string); !ok || len(str) == 0 {
					return errors.New("Branch name can't be empty")
				}

				return nil
			},
		},
	}

	return branchNameMessage
}
