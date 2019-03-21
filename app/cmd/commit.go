package root

import (
	"fmt"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes",
	Run: func(cmd *cobra.Command, args []string) {

		untrackedFilesCommander := run.Commander{
			Command: "git",
			Args:    []string{"status"},
		}

		x, err := git.Status(untrackedFilesCommander)

		untrackedFiles := git.GetFiles(x)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		var selectedFiles []string
		if len(untrackedFiles) > 0 {
			//Get user to select files to commit
			selectedFiles = selectFilesToTrack(untrackedFiles)
		}

		fmt.Println(selectedFiles)

		commitCommander := run.Commander{
			Command: "git",
			Args:    []string{"commit", "-am", "test commit"},
		}

		result, err := git.CommitChanges(commitCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println(result)
	},
}

func selectFilesToTrack(availableFiles []string) []string {
	var result string
	var err error
	var exit = false

	prompt := promptui.SelectWithAdd{
		Label: "What's your text editor",
		Items: availableFiles,
	}

	for !exit {
		index := -1
		for index < 0 {

			index, result, err = prompt.Run()
			fmt.Println(index)

			availableFiles = append(availableFiles[:index], availableFiles[index+1:]...)

			if result == "Exit" {
				exit = true
			}
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return availableFiles
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
