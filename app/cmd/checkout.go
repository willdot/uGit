package root

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/app/pkg/git"
	"github.com/willdot/uGit/app/pkg/input"
	"github.com/willdot/uGit/app/pkg/run"

	tea "github.com/charmbracelet/bubbletea"
)

const exit = "Exit"

var newBranchFlag bool

var checkoutCmd = &cobra.Command{
	Use:   "cko",
	Short: "Checkout a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if newBranchFlag {
			p := tea.NewProgram(input.InitialTextInputModel("enter branch name"))
			model, err := p.Run()
			if err != nil {
				log.Fatal(err)
			}

			m, ok := model.(input.TextInputModel)
			if !ok {
				return
			}
			if m.Err != nil {
				fmt.Printf("error: %s\n", m.Err)
				return
			}

			branchName := m.TextInput.Value()

			if branchName == "" {
				os.Exit(1)
			}

			checkout(branchName, newBranchFlag)
			return
		}

		branchCommander := run.Commander{
			Command: "git",
			Args:    []string{"branch", "-a"},
		}

		branches, err := git.GetBranches(branchCommander)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		branchName := askUserToSelectSingleOption(branches)
		if branchName == "" {
			return
		}

		checkout(branchName, newBranchFlag)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolVarP(&newBranchFlag, "new", "n", false, "create new branch")
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
