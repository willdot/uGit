package git

import (
	"fmt"
	"strings"
	"uGit/app/pkg/run"
)

// Status will run git status and return the result
func Status(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}

// GetFiles will return a slice of files that aren't tracked
func GetFiles(s string) []string {

	var untracked bool

	x := strings.Split(s, "\t")

	var result []string

	for i := 0; i < len(x); i++ {
		line := strings.Trim(x[i], " ")
		line = strings.TrimSuffix(line, "\n")
		fmt.Println(line + ".")

		if line == "Untracked files:" {
			untracked = true
			break
		}

		if line == "nothing to commit, working tree clean" {
			return nil
		}
	}
	if untracked {
		result = getUntracked(x)
	}

	return result
}

func getUntracked(s []string) []string {

	var result []string

	for i := 0; i < len(s); i++ {
		line := strings.Trim(s[i], " ")
		line = strings.TrimSuffix(line, "\n")

		if line == "Untracked files:" {
			for x := i + 1; x < len(s); x++ {
				line = strings.Trim(s[x], " ")
				line = strings.TrimSuffix(line, "\n\n")
				line = strings.TrimSuffix(line, "\n")

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
