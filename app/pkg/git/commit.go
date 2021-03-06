package git

import "github.com/willdot/uGit/app/pkg/run"

// CommitChanges will run git commit
func CommitChanges(commander run.ICommander) (string, error) {
	result, err := commander.CommandWithResult()

	return string(result), err
}
