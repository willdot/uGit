package git

import "uGit/app/pkg/run"

// DeleteBranch will run git delete for a branch
func DeleteBranch(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}
