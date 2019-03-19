package root

import (
	"github.com/spf13/cobra"
)

// checkoutCmd represents the say command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {

		/*branchCommander := run.RealCommander{
			command: "git",
			args:    []string{"branch"},
		}
		result, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		branches := git.SplitBranches(result)

		branches = git.RemoveCurrentBranch(branches)

		prompt := promptui.Select{
			Label:    "Select branch",
			Items:    branches,
			HideHelp: true,
		}

		_, selection, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		checkoutCommander := run.RealCommander{
			command: "git",
			args:    []string{"checkout", strings.Replace(selection, " ", "", -1)},
		}

		x, err := git.CheckoutBranch(commander)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
		}

		fmt.Println(x)*/
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
