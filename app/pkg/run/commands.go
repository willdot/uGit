package run

import (
	"os/exec"
)

// Commander is an interface for running os/exec Commands
type Commander interface {
	CombinedOutput(string, ...string) ([]byte, error)
}

// RealCommander is a real struct that can be used
type RealCommander struct{}

func (r RealCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

// CommandWithResult will run a command and return the output or an error
func CommandWithResult(commander Commander, command string, args ...string) (string, error) {

	output, err := commander.CombinedOutput(command, args...)

	return string(output), err
}

// CommandWithoutResult will run a command and only return an error if one is found
func CommandWithoutResult(commander Commander, command string, args ...string) error {
	_, err := CommandWithResult(commander, command, args...)

	return err
}
