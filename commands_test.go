package main

import "testing"

func TestCommandWithResultReturnsAStringValue(t *testing.T) {
	dontWant := ""
	got, _ := RunCommandWithResult("ls")

	if dontWant == got {
		t.Errorf("didn't want nothing, but got '%s'", got)
	}
}

func TestCommandWithResultInvalidCommandReturnsError(t *testing.T) {

	_, err := RunCommandWithResult("hh")

	if err == nil {
		t.Errorf("wanted error but didn't get one")
	}
}
