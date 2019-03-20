package run

import (
	"os/exec"
)

// Commander is an interface for running os/exec Commands
type Commander interface {
	CombinedOutput() ([]byte, error)
}

// RealCommander is a real struct that can be used
type RealCommander struct {
	Command string
	Args    []string
}

// CombinedOutput runs an os command and returns the result
func (r RealCommander) CombinedOutput() ([]byte, error) {
	return exec.Command(r.Command, r.Args...).CombinedOutput()
}

// CommandWithResult will run a command and return the output or an error
func CommandWithResult(commander Commander) (string, error) {

	output, err := commander.CombinedOutput()

	return string(output), err
}
