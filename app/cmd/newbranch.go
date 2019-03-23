package root

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// checkoutNewCmd represents the checkout of a new branch command
var checkoutNewCmd = &cobra.Command{
	Use:   "checkout-new",
	Short: "Checkout a new branch",
	Run: func(cmd *cobra.Command, args []string) {

		var selection string

		question := getBranchNameQuestion()
		err := survey.Ask(question, &selection)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		checkout(selection, true)
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
				Message: "Enter a branch name",
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
