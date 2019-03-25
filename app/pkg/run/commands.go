package run

import (
	"os/exec"
)

// ICommander is an interface for running os/exec Commands
type ICommander interface {
	//RunCommand() ([]byte, error)
	CommandWithResult() (string, error)
}

// Commander is a real struct that can be used to run os commands.
// Command is the original command
// Args is any other commands / arguments
type Commander struct {
	Command string
	Args    []string
}

// runCommand runs an os command and returns the result
func runCommand(command string, args []string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

// CommandWithResult will run a command and return the output or an error
func (r Commander) CommandWithResult() (string, error) {

	output, err := runCommand(r.Command, r.Args)

	return string(output), err
}
