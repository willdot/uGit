package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/run"
)

func MergeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "merge",
		Short: "Merge branches",
		Run: func(cmd *cobra.Command, args []string) {
			err := merge()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd
}

func merge() error {
	branches, err := git.GetBranches()
	if err != nil {
		return err
	}

	branchName, err := askUserToSelectSingleOption(branches, "")
	if err != nil {
		return err
	}

	if branchName == "" {
		return nil
	}

	branchName = git.RemoveRemoteOriginFromName(branchName)

	args := []string{"merge", strings.TrimSpace(branchName)}

	result, err := run.RunCommand("git", args)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
