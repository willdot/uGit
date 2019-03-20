package run

import (
	"testing"
)

func TestCommandWithResultReturnsAStringValue(t *testing.T) {

	t.Run("run and returns a string", func(t *testing.T) {
		dontWant := ""
		commander := Commander{
			Command: "go",
			Args:    []string{"env"},
		}
		got, _ := CommandWithResult(commander)

		if dontWant == got {
			t.Errorf("didn't want nothing, but got '%s'", got)
		}
	})

	t.Run("run with with error returned", func(t *testing.T) {
		commander := Commander{
			Command: "hh",
		}
		_, err := CommandWithResult(commander)

		if err == nil {
			t.Errorf("wanted error but didn't get one")
		}
	})
}
