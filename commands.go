package main

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

type fakeCommander struct {
	result []byte
	err    error
}

func (f fakeCommander) combinedOutput(command string, args ...string) ([]byte, error) {

	if f.err != nil {
		return nil, f.err
	}

	return f.result, nil
}

// RunCommandWithResult will run a command and return the output or an error
func RunCommandWithResult(commander Commander, command string, args ...string) (string, error) {

	output, err := commander.combinedOutput(command, args...)

	return string(output), err
}

// RunCommandWithoutResult will run a command and only return an error if one is found
func RunCommandWithoutResult(commander Commander, command string, args ...string) error {
	_, err := RunCommandWithResult(commander, command, args...)

	return err
}
