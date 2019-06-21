package git

import "uGit/app/pkg/run"

// Merge will run git merge and return the result
func Merge(commander run.ICommander) (string, error) {
	result, err := commander.CommandWithResult()

	return string(result), err
}
