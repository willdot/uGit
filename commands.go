package main

import (
	"os/exec"
)

// RunCommandWithResult will run a command and return the output or an error
func RunCommandWithResult(command string, args ...string) (string, error) {

	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()

	return string(output), err
}

// RunCommandWithoutResult will run a command and only return an error if one is found
func RunCommandWithoutResult(command string, args ...string) error {
	_, err := RunCommandWithResult(command, args...)

	return err
}
