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
