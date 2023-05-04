package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/input"
	"github.com/willdot/uGit/pkg/run"

	tea "github.com/charmbracelet/bubbletea"
)

const exit = "Exit"

var newBranchFlag bool

func CheckoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cko",
		Short: "Checkout a branch",
		Run: func(cmd *cobra.Command, args []string) {
			err := checkout()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&newBranchFlag, "new", "n", false, "create new branch")

	return cmd
}
func checkout() error {
	if newBranchFlag {
		return checkoutNewBranch()
	}

	branches, err := git.GetBranches()
	if err != nil {
		return err
	}

	branchName := askUserToSelectSingleOption(branches, "")
	if branchName == "" {
		return nil
	}

	performCheckout(branchName, newBranchFlag)

	return nil
}

func checkoutNewBranch() error {
	p := tea.NewProgram(input.InitialTextInputModel("enter branch name"))
	model, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	m, ok := model.(input.TextInputModel)
	if !ok {
		return nil
	}
	if m.Err != nil {
		return m.Err
	}

	branchName := m.TextInput.Value()

	if branchName == "" {
		return nil
	}

	performCheckout(branchName, newBranchFlag)
	return nil
}

func performCheckout(branchSelection string, new bool) {

	branchSelection = git.RemoveRemoteOriginFromName(branchSelection)

	args := []string{"checkout"}

	if new {
		args = append(args, "-b")
	}

	args = append(args, strings.TrimSpace(branchSelection))

	result, err := run.RunCommand("git", args)
	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
	}

	fmt.Println(result)
}
