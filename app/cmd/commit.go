package root

import (
	"fmt"
	"uGit/app/pkg/git"
	"uGit/app/pkg/run"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
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

		untrackedFiles, nothingToCommit := git.GetFiles(x)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		if nothingToCommit {
			fmt.Println("Nothing to commit")
			return
		}

		if len(untrackedFiles) > 0 {
			var selectedFiles []string

			selectedFiles = selectFilesToTrack(untrackedFiles)
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

		commitMessage := ""

		commitQ := getCommitQuestion()
		err = survey.Ask(commitQ, &commitMessage)

		if commitMessage == "exit" {
			return
		}

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		commitCommander := run.Commander{
			Command: "git",
			Args:    []string{"commit", "-am", commitMessage},
		}

		commitResult, err := git.CommitChanges(commitCommander)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			return
		}

		fmt.Println(commitResult)
	},
}

func getCommitQuestion() []*survey.Question {
	var commitMessage = []*survey.Question{
		{
			Name: "commit",
			Prompt: &survey.Input{
				Message: `Enter a commit message or type "exit" to cancel`,
			},
			Validate: survey.Required,
		},
	}

	return commitMessage
}

func selectFilesToTrack(availableFiles []string) []string {

	options := append([]string{"**Select all**", "**Exit and ignore selections**"}, availableFiles...)

	result := []string{}
	prompt := &survey.MultiSelect{
		Message: "You have untracked files. Select files to add.",
		Options: options,
	}

	survey.AskOne(prompt, &result, nil)

	for i := 0; i < len(result); i++ {
		if result[i] == "**Select all**" {
			return availableFiles
		}
		if result[i] == "**Exit and ignore selections**" {
			return nil
		}
	}

	return result
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
