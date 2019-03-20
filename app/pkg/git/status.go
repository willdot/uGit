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

// GetUntrackedFiles will return a slice of files that aren't tracked
func GetUntrackedFiles(s string) []string {

	x := strings.Split(s, `(use "git add <file>..." to include in what will be committed)`)[1]
	x = strings.Split(x, "nothing added")[0]

	untrackedSection := strings.Split(x, "\n")
	var result []string
	for i := 0; i < len(untrackedSection); i++ {
		if untrackedSection[i] != "\t" && untrackedSection[i] != "" {
			file := strings.Replace(untrackedSection[i], "\t", "", -1)
			file = strings.Trim(file, " ")
			result = append(result, file)
		}
	}

	return result
}
