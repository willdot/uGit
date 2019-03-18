package main

import "testing"

func TestCommandWithResultReturnsAStringValue(t *testing.T) {

	t.Run("run and returns a string", func(t *testing.T) {
		dontWant := ""
		got, _ := RunCommandWithResult("go", "env")

		if dontWant == got {
			t.Errorf("didn't want nothing, but got '%s'", got)
		}
	})

	t.Run("run with with error returned", func(t *testing.T) {
		_, err := RunCommandWithResult("hh")

		if err == nil {
			t.Errorf("wanted error but didn't get one")
		}
	})
}

func TestCommandWithoutResult(t *testing.T) {

	t.Run("run with no errors", func(t *testing.T) {
		err := RunCommandWithoutResult("go", "env")

		if err != nil {
			t.Errorf("didn't want an error but got one '%s'", err)
		}
	})

	t.Run("run with with error returned", func(t *testing.T) {
		err := RunCommandWithoutResult("hh")

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
