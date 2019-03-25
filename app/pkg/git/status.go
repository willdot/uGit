package git

import (
	"strings"
	"uGit/app/pkg/run"
)

// Status will run git status and return the result
func Status(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}

// GetFilesOrNothingToCommit will return a slice of files that aren't tracked and if there is nothing to commit, a true flag
func GetFilesOrNothingToCommit(s string) ([]string, bool) {

	x := strings.Split(s, "\n")

	var result []string

	for i := 0; i < len(x); i++ {
		line := strings.TrimSpace(x[i])

		if line == "Untracked files:" {
			x = x[i:]
			result = getUntracked(x)
			break
		}

		if line == "nothing to commit, working tree clean" {
			return nil, true
		}
	}

	return result, false
}

func getUntracked(s []string) []string {

	var result []string

	for i := 0; i < len(s); i++ {
		line := strings.TrimSpace(s[i])

		if line == "Untracked files:" {
			for x := i + 1; x < len(s); x++ {
				line = strings.TrimSpace(s[x])

				if strings.HasPrefix(line, "no changes added to commit") || strings.HasPrefix(line, "nothing added to commit ") {
					break
				}

				if !strings.HasPrefix(line, `(use "git add <file>..."`) && line != "" {
					result = append(result, line)
				}
			}
		}
	}

	return result
}

// GetNotStagedFiles will return a slice of files that are not staged for commit
func GetNotStagedFiles(s string) []string {
	splitLines := strings.Split(s, "\n")

	var result []string

	for i := 0; i < len(splitLines); i++ {
		line := strings.TrimSpace(splitLines[i])

		// We want to get the files after this line
		if line == "Changes not staged for commit:" {
			notStagedLines := splitLines[i:]
			// There is an initial blank line before the list of files not staged.
			// So that we know we have the blank file after the list of files not tracked, use this flag
			initialBlankLine := false

			for j := 3; j < len(notStagedLines); j++ {
				line = strings.TrimSpace(notStagedLines[j])

				if line == "" {
					if initialBlankLine == false {
						initialBlankLine = true
					} else {
						// No more files that are not staged
						break
					}
				} else {
					result = append(result, line)
				}
			}
			break
		}
	}

	return result
}
