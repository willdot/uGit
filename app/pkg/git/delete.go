package git

import "github.com/willdot/uGit/app/pkg/run"

// DeleteBranch will run git delete for a branch
func DeleteBranch(commander run.ICommander) (string, error) {
	result, err := commander.CommandWithResult()

	return string(result), err
}
