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
	command string
	args    []string
}

// CombinedOutput runs an os command and returns the result
func (r RealCommander) CombinedOutput() ([]byte, error) {
	return exec.Command(r.command, r.args...).CombinedOutput()
}

// CommandWithResult will run a command and return the output or an error
func CommandWithResult(commander Commander) (string, error) {

	output, err := commander.CombinedOutput()

	return string(output), err
}

// CommandWithoutResult will run a command and only return an error if one is found
/*func CommandWithoutResult(commander Commander) error {
	_, err := CommandWithResult(commander)

	return err
}*/
