package root

import (
	"fmt"
	"strings"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout a branch command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {

		branchCommander := run.Commander{
			Command: "git",
			Args:    []string{"branch"},
		}
		branches, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		branchesSlice := git.SplitBranches(branches, true)

		prompt := promptui.Select{
			Label:    "Select branch",
			Items:    branchesSlice,
			HideHelp: true,
		}

		_, selection, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		checkout(selection)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func checkout(branchSelection string) {
	checkoutCommander := run.Commander{
		Command: "git",
		Args:    []string{"checkout"},
	}

	result, err := git.CheckoutBranch(checkoutCommander, strings.Replace(branchSelection, " ", "", -1))

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}
