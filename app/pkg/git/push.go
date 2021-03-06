package git

import "github.com/willdot/uGit/app/pkg/run"

// Push will run a git push command and return the result
func Push(commander run.ICommander) (string, error) {
	result, err := commander.CommandWithResult()

	return string(result), err
}
