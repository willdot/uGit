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

var pushFlag bool

var commitCmd = &cobra.Command{
	Use:   "com",
	Short: "Commit changes",
	Run: func(cmd *cobra.Command, args []string) {

		workOutFilesToBeCommitted()

		status := getStatus()

		filesToBeCommitted := git.GetFilesToBeCommitted(status)

		if len(filesToBeCommitted) == 0 {
			fmt.Println("Nothing to commit")
			return
		}

		fmt.Println("Files to be committed")
		for _, file := range filesToBeCommitted {
			fmt.Println(file)
		}

		commit()

		if pushFlag {
			push()
		}

	},
}

func workOutFilesToBeCommitted() {
	status := getStatus()

	untrackedFiles, nothingToCommit := git.GetFilesOrNothingToCommit(status)

	if nothingToCommit {
		fmt.Println("Nothing to commit")
		os.Exit(1)
	}

	if len(untrackedFiles) > 0 {
		resolveUntrackedFiles(untrackedFiles)
	}

	notStaged := git.GetNotStagedFiles(status)

	if len(notStaged) > 0 {
		stageFiles(notStaged)
	}
}

func getStatus() string {
	untrackedFilesCommander := run.Commander{
		Command: "git",
		Args:    []string{"status"},
	}

	status, err := git.Status(untrackedFilesCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	return status
}

func resolveUntrackedFiles(untrackedFiles []string) {
	var selectedFiles []string

	selectedFiles = askUserToSelectOptions(untrackedFiles, "You have untracked files. Select files to add.", true)

	if len(selectedFiles) > 0 {
		addFiles(selectedFiles)
	}
}

func stageFiles(availableFiles []string) {
	selectedFiles := askUserToSelectOptions(availableFiles, "You have unstaged files. Select files to add.", true)

	if len(selectedFiles) > 0 {
		var filesToAdd []string

		for _, file := range selectedFiles {
			file = strings.Split(file, ":")[1]
			filesToAdd = append(filesToAdd, strings.TrimSpace(file))
		}

		addFiles(filesToAdd)
	}
}

func printSelectedFiles(files []string) {
	fmt.Println("You selected:")

	for _, file := range files {
		fmt.Println(file)
	}
}

func addFiles(filesToAdd []string) {
	if len(filesToAdd) > 0 {

		printSelectedFiles(filesToAdd)

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

func push() {
	pushCommander := run.Commander{
		Command: "git",
		Args:    []string{"push"},
	}

	result, err := run.CommandWithResult(pushCommander)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	fmt.Println(result)
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
	commitCmd.Flags().BoolVarP(&pushFlag, "push", "p", false, "push after commit")
}
