package run

import (
	"os/exec"
)

// ICommander is an interface for running os/exec Commands
type ICommander interface {
	RunCommand() ([]byte, error)
}

// Commander is a real struct that can be used to run os commands.
// Command is the original command
// Args is any other commands / arguments
type Commander struct {
	Command string
	Args    []string
}

// RunCommand runs an os command and returns the result
func (r Commander) RunCommand() ([]byte, error) {
	return exec.Command(r.Command, r.Args...).CombinedOutput()
}

// CommandWithResult will run a command and return the output or an error
func CommandWithResult(commander ICommander) (string, error) {

	output, err := commander.RunCommand()

	return string(output), err
}
