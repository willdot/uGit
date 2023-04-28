package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/git"
	"github.com/willdot/uGit/pkg/input"
	"github.com/willdot/uGit/pkg/run"
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
	// untrackedFilesCommander := run.Commander{
	// 	Command: "git",
	// 	Args:    []string{"status"},
	// }

	// status, err := git.Status(untrackedFilesCommander)

	status, err := run.RunCommand("git", []string{"status"})

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	return status
}

func resolveUntrackedFiles(untrackedFiles []string) {
	var selectedFiles []string

	selectedFiles = askUserToSelectOptions(untrackedFiles, "Untracked files. Select files to add.", true)

	if len(selectedFiles) > 0 {
		addFiles(selectedFiles)
	}
}

func stageFiles(availableFiles []string) {
	selectedFiles := askUserToSelectOptions(availableFiles, "Unstaged files. Select files to add.", true)

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

		// addFilesCommander := run.Commander{
		// 	Command: "git",
		// 	Args:    append([]string{"add"}, filesToAdd...),
		// }

		// _, err := git.Add(addFilesCommander)

		_, err := run.RunCommand("git", append([]string{"add"}, filesToAdd...))

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
	}
}

func commit() {
	p := tea.NewProgram(input.InitialTextInputModel("enter commit message"))
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

	commitMessage := m.TextInput.Value()
	if commitMessage == "" {
		return
	}

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	// commitCommander := run.Commander{
	// 	Command: "git",
	// 	Args:    []string{"commit", "-m", commitMessage},
	// }

	// commitResult, err := git.CommitChanges(commitCommander)

	commitResult, err := run.RunCommand("git", []string{"commit", "-m", commitMessage})

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		os.Exit(1)
	}

	fmt.Println(commitResult)
}

func push() {
	// pushCommander := run.Commander{
	// 	Command: "git",
	// 	Args:    []string{"push"},
	// }

	// result, err := git.Push(pushCommander)

	result, err := run.RunCommand("git", []string{"push"})

	if err != nil {
		handleErrorPush(result)
	} else {
		fmt.Println(result)
	}
}

func pushSetUpstream(command string) {

	args := strings.Split(strings.TrimSpace(command), " ")
	// pushCommander := run.Commander{
	// 	Command: "git",
	// 	Args:    args[1:],
	// }

	// result, _ := git.Push(pushCommander)
	result, err := run.RunCommand("git", args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func handleErrorPush(errorMessage string) {
	lines := strings.Split(errorMessage, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "To push the current branch and set the remote as upstream, use") {
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

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().BoolVarP(&pushFlag, "push", "p", false, "push after commit")
}
