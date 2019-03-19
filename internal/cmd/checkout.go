package root

import (
	"fmt"
	"strings"
	"uGit/internal/pkg/git"
	"uGit/internal/pkg/run"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the say command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Select a branch")

		commander := run.RealCommander{}
		result, err := git.GetBranches(commander)

		branches := git.SplitBranches(result)

		prompt := promptui.Select{
			Label: "Select branch",
			Items: branches,
		}

		_, selection, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		x, err := git.CheckoutBranch(commander, strings.Replace(selection, " ", "", -1))

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
		}

		fmt.Println(x)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
