package git

import "uGit/app/pkg/run"

// Push will run a git push command and return the result
func Push(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}
