package cli

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/input"
	"github.com/willdot/uGit/pkg/run"
)

var pushFlag bool

func CommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "com",
		Short: "Commit changes",
		Run: func(cmd *cobra.Command, args []string) {
			err := commit()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&pushFlag, "push", "p", false, "push after commit")

	return cmd
}

func commit() error {
	filesToBeCommitted, err := workOutFilesToBeCommitted()
	if err != nil {
		return err
	}
	if len(filesToBeCommitted) == 0 {
		fmt.Println("Nothing to commit")
		return nil
	}

	fmt.Println("Files to be committed")
	for _, file := range filesToBeCommitted {
		fmt.Println(file)
	}

	err = makeCommit()
	if err != nil {
		return err
	}

	if pushFlag {
		push()
	}

	return nil
}

func workOutFilesToBeCommitted() ([]string, error) {
	status, err := getStatus()
	if err != nil {
		return nil, err
	}

	untrackedFiles, nothingToCommit := git.GetFilesOrNothingToCommit(status)

	if nothingToCommit {
		return nil, nil
	}

	if len(untrackedFiles) > 0 {
		resolveUntrackedFiles(untrackedFiles)
	}

	notStaged := git.GetNotStagedFiles(status)

	if len(notStaged) > 0 {
		err := stageFiles(notStaged)
		if err != nil {
			return nil, err
		}
	}

	status, err = getStatus()
	if err != nil {
		return nil, err
	}

	return git.GetFilesToBeCommitted(status), nil
}

func getStatus() (string, error) {
	status, err := run.RunCommand("git", []string{"status"})

	if err != nil {
		return "", err
	}

	return status, nil
}

func resolveUntrackedFiles(untrackedFiles []string) {
	var selectedFiles []string

	selectedFiles = askUserToSelectOptions(untrackedFiles, "Untracked files. Select files to add.", true)

	if len(selectedFiles) > 0 {
		addFiles(selectedFiles)
	}
}

func stageFiles(availableFiles []string) error {
	selectedFiles := askUserToSelectOptions(availableFiles, "Unstaged files. Select files to add.", true)

	if len(selectedFiles) > 0 {
		var filesToAdd []string

		for _, file := range selectedFiles {
			file = strings.Split(file, ":")[1]
			filesToAdd = append(filesToAdd, strings.TrimSpace(file))
		}

		err := addFiles(filesToAdd)
		if err != nil {
			return err
		}
	}

	return nil
}

func printSelectedFiles(files []string) {
	fmt.Println("You selected:")

	for _, file := range files {
		fmt.Println(file)
	}
}

func addFiles(filesToAdd []string) error {
	if len(filesToAdd) > 0 {

		printSelectedFiles(filesToAdd)

		_, err := run.RunCommand("git", append([]string{"add"}, filesToAdd...))
		if err != nil {
			return err
		}
	}

	return nil
}

func makeCommit() error {
	p := tea.NewProgram(input.InitialTextInputModel("enter commit message"))
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

	commitMessage := m.TextInput.Value()
	if commitMessage == "" {
		return nil
	}

	if err != nil {
		return err
	}

	commitResult, err := run.RunCommand("git", []string{"commit", "-m", commitMessage})
	if err != nil {
		return err
	}

	fmt.Println(commitResult)

	return nil
}

func push() {
	result, err := run.RunCommand("git", []string{"push"})

	if err == nil {
		fmt.Println(result)
		return
	}

	handleErrorPush(err)
}

func pushSetUpstream(command string) {

	args := strings.Split(strings.TrimSpace(command), " ")

	result, err := run.RunCommand("git", args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func handleErrorPush(err error) {
	lines := strings.Split(err.Error(), "\n")

	for _, line := range lines {
		if strings.Contains(line, "To push the current branch and set the remote as upstream, use") {
			fmt.Println(lines[0])

			result := false

			res := askUserToSelectSingleOption([]string{"yes", "no"}, "Would you like to set remote as upstream?")
			if res == "yes" {
				result = true
			}

			if !result {
				return
			}
		}

		if strings.HasPrefix(strings.TrimSpace(line), "git push") {
			pushSetUpstream(line)
		}
	}
}
