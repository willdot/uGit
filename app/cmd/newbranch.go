package root

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// checkoutNewCmd represents the checkout of a new branch command
var checkoutNewCmd = &cobra.Command{
	Use:   "checkout new",
	Short: "Checkout a new branch",
	Run: func(cmd *cobra.Command, args []string) {

		var selection string

		question := getBranchNameQuestion()
		err := survey.Ask(question, &selection)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		checkout(selection, true)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func getBranchNameQuestion() []*survey.Question {
	var commitMessage = []*survey.Question{
		{
			Name: "commit",
			Prompt: &survey.Input{
				Message: "Enter a commit message",
			},
			Validate: func(val interface{}) error {
				// since we are validating an Input, the assertion will always succeed

				x := strings.Replace(val.(string), " ", "", -1)

				if str, ok := val.(string); !ok || len(str) != len(x) {
					return errors.New("Branch name cannot contain spaces")
				}
				return nil
			},
		},
	}

	return commitMessage
}
