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
			Args:    []string{"branch", "-a"},
		}
		branches, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		branchesSlice := git.SplitBranches(branches, true)

		help := promptui.Styler(promptui.FGRed)("Use arrow keys (or J K) to go up and down.")
		searchHelp := promptui.Styler(promptui.FGYellow)("Use ? to toggle search.")
		fmt.Println(help)
		fmt.Println(searchHelp)

		searcher := func(input string, index int) bool {
			b := branchesSlice[index]
			branchName := strings.Replace(strings.ToLower(b), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(branchName, input)
		}

		searchKey := promptui.Key{
			Code:    63,
			Display: "?",
		}

		selectKeys := &promptui.SelectKeys{
			Search: searchKey,
		}

		iconSelect := promptui.Styler(promptui.FGBlue)("*")
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   iconSelect + "{{. | blue }}",
			Inactive: "  {{ . }}",
		}

		prompt := promptui.Select{
			Label:     "Select branch",
			Items:     branchesSlice,
			HideHelp:  true,
			Searcher:  searcher,
			Keys:      selectKeys,
			Templates: templates,
		}

		_, selection, err := prompt.Run()

		fmt.Println(selection)

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
