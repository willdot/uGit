package git

import "uGit/app/pkg/run"

// Add will run git status and return the result
func Add(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}
