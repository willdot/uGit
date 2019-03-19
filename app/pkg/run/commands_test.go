package run

import (
	"testing"
)

func TestCommandWithResultReturnsAStringValue(t *testing.T) {

	t.Run("run and returns a string", func(t *testing.T) {
		dontWant := ""
		commander := RealCommander{
			command: "go",
			args:    []string{"env"},
		}
		got, _ := CommandWithResult(commander)

		if dontWant == got {
			t.Errorf("didn't want nothing, but got '%s'", got)
		}
	})

	t.Run("run with with error returned", func(t *testing.T) {
		commander := RealCommander{
			command: "hh",
		}
		_, err := CommandWithResult(commander)

		if err == nil {
			t.Errorf("wanted error but didn't get one")
		}
	})
}

/*func TestCommandWithoutResult(t *testing.T) {

	commander := RealCommander{}
	t.Run("run with no errors", func(t *testing.T) {
		err := CommandWithoutResult(commander, "go", "env")

		if err != nil {
			t.Errorf("didn't want an error but got one '%s'", err)
		}
	})

	t.Run("run with with error returned", func(t *testing.T) {
		err := CommandWithoutResult(commander, "hh")

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}*/
