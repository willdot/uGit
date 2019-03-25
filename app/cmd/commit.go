package root

import (
	"fmt"
	"os"
	"strings"
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

		status, err := git.Status(untrackedFilesCommander)

		untrackedFiles, nothingToCommit := git.GetFilesOrNothingToCommit(status)

		if err != nil {
			fmt.Printf("error: %v", errors.WithMessage(err, ""))
			os.Exit(1)
		}

		if nothingToCommit {
			fmt.Println("Nothing to commit")
			return
		}

		if len(untrackedFiles) > 0 {
			resolveUntrackedFiles(untrackedFiles)
		}

		notStaged := git.GetNotStagedFiles(status)

		if len(notStaged) > 0 {
			stageFiles(notStaged)
		}

		status, err = git.Status(untrackedFilesCommander)

		filesToBeCommitted := git.GetFilesToBeCommitted(status)

		fmt.Println("Files to be committed")
		for _, file := range filesToBeCommitted {
			fmt.Println(file)
		}

		commit()
	},
}

func resolveUntrackedFiles(untrackedFiles []string) {
	var selectedFiles []string

	selectedFiles = selectFilesToTrack(untrackedFiles)

	if len(selectedFiles) > 0 {

		printSelectedFiles(selectedFiles)

		addFilesCommander := run.Commander{
			Command: "git",
			Args:    append([]string{"add"}, selectedFiles...),
		}

		_, err := git.Add(addFilesCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	}
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

func stageFiles(availableFiles []string) {
	selectedFiles := selectFilesToStage(availableFiles)

	if len(selectedFiles) > 0 {
		var filesToAdd []string

		for _, file := range selectedFiles {
			file = strings.Split(file, ":")[1]
			filesToAdd = append(filesToAdd, strings.TrimSpace(file))
		}

		addFilesCommander := run.Commander{
			Command: "git",
			Args:    append([]string{"add"}, filesToAdd...),
		}

		_, err := git.Add(addFilesCommander)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	}
}

func selectFilesToStage(availableFiles []string) []string {
	options := append([]string{"**Select all**", "**Exit and ignore selections**"}, availableFiles...)

	result := []string{}
	prompt := &survey.MultiSelect{
		Message: "You have unstaged files. Select files to add.",
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

func printSelectedFiles(files []string) {
	fmt.Println("You selected:")

	for _, file := range files {
		fmt.Println(file)
	}
}

func commit() {
	commitMessage := ""

	commitQ := getCommitQuestion()
	err := survey.Ask(commitQ, &commitMessage)

	if commitMessage == "exit" {
		return
	}

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	commitCommander := run.Commander{
		Command: "git",
		Args:    []string{"commit", "-m", commitMessage},
	}

	commitResult, err := git.CommitChanges(commitCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	fmt.Println(commitResult)
}

func getCommitQuestion() []*survey.Question {
	var commitQuestion = []*survey.Question{
		{
			Name: "commit",
			Prompt: &survey.Input{
				Message: `Enter a commit message or type "exit" to cancel`,
			},
			Validate: survey.Required,
		},
	}

	return commitQuestion
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
