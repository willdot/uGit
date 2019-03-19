package run

import (
	"os/exec"
)

// Commander is an interface for running os/exec Commands
type Commander interface {
	combinedOutput(string, ...string) ([]byte, error)
}

// RealCommander is a real struct that can be used
type RealCommander struct{}

func (r RealCommander) combinedOutput(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

// FakeCommander is used for mocking
type FakeCommander struct {
	Result []byte
	Err    error
}

func (f FakeCommander) combinedOutput(command string, args ...string) ([]byte, error) {

	if f.Err != nil {
		return nil, f.Err
	}

	return f.Result, nil
}

// CommandWithResult will run a command and return the output or an error
func CommandWithResult(commander Commander, command string, args ...string) (string, error) {

	output, err := commander.combinedOutput(command, args...)

	return string(output), err
}

// CommandWithoutResult will run a command and only return an error if one is found
func CommandWithoutResult(commander Commander, command string, args ...string) error {
	_, err := CommandWithResult(commander, command, args...)

	return err
}
