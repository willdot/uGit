package run

import (
	"os/exec"

	"github.com/pkg/errors"
)

func RunCommand(command string, args []string) (string, error) {
	result, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(result))
	}

	return string(result), nil
}
