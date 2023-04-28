package run

import (
	"os/exec"
)

func RunCommand(command string, args []string) (string, error) {
	result, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(result), nil
}