package root

import (
	"fmt"
	"sync"
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

		if len(untrackedFiles) > 0 {
			var selectedFiles []string

			var waitGroup sync.WaitGroup

			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				fmt.Println("You have untracked files. Please select any files you wish to add")
				//Get user to select files to commit
				selectedFiles = selectFilesToTrack(untrackedFiles)

			}()

			waitGroup.Wait()
			fmt.Println("You selected")
			fmt.Println(selectedFiles)

			if len(selectedFiles) > 0 {
				addFilesCommander := run.Commander{
					Command: "git",
					Args:    append([]string{"add"}, selectedFiles...),
				}

				x, err := git.Add(addFilesCommander)

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				fmt.Println(x)
			}
		}

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

	options := append([]string{"**Select all**", "**Exit**", "**Exit and ignore selections**"}, availableFiles...)
	var result string
	var err error
	var exit = false

	var selectedFiles []string

	for !exit {
		prompt := promptui.Select{
			Label: "What's your text editor",
			Items: options,
		}
		index := -1
		for index < 0 {

			index, result, err = prompt.Run()

			switch result {
			case "**Exit**":
				exit = true
			case "**Exit and ignore selections**":
				exit = true
				selectedFiles = nil
			case "**Select all**":
				// If all files have already been selected by user,do nothing
				if len(options) > 3 {
					selectedFiles = availableFiles
					options = append(options[:3])
				}
			default:
				selectedFiles = append(selectedFiles, result)
				options = append(options[:index], options[index+1:]...)
			}
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return selectedFiles
}

func addAllFiles(s []string) []string {

	var result []string
	for i := 0; i < len(s); i++ {
		file := s[i]
		result = append(result, file)
	}
	return result
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
