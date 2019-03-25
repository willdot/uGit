package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// checkoutNewCmd represents the checkout of a new branch command
var checkoutNewCmd = &cobra.Command{
	Use:   "cko-n",
	Short: "Checkout a new branch",
	Run: func(cmd *cobra.Command, args []string) {

		var branchName string

		question := getBranchNameQuestion()
		err := survey.Ask(question, &branchName)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		if branchName == "exit" {
			os.Exit(1)
		}

		checkout(branchName, true)
	},
}

func init() {
	rootCmd.AddCommand(checkoutNewCmd)
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
