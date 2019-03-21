package git

import "uGit/app/pkg/run"

// CommitChanges will run git commit
func CommitChanges(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return string(result), err
}
