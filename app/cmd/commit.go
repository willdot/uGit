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
			fmt.Println("You have untracked files. Please select any files you wish to add")
			//Get user to select files to commit
			selectedFiles = selectFilesToTrack(untrackedFiles)
		}

		fmt.Println(selectedFiles)

		/*commitCommander := run.Commander{
			Command: "git",
			Args:    []string{"commit", "-am", "test commit"},
		}

		result, err := git.CommitChanges(commitCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println(result)*/
	},
}

func selectFilesToTrack(availableFiles []string) []string {

	availableFiles = append(availableFiles, "Exit")
	var result string
	var err error
	var exit = false

	var selectedFiles []string

	for !exit {
		prompt := promptui.Select{
			Label: "What's your text editor",
			Items: availableFiles,
		}
		index := -1
		for index < 0 {

			index, result, err = prompt.Run()

			if result == "Exit" {
				exit = true
			} else {
				availableFiles = append(availableFiles[:index], availableFiles[index+1:]...)
				selectedFiles = append(selectedFiles, result)
			}
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return selectedFiles
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
