package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitBranch(t *testing.T) {
	want := []string{"Dev", "Master"}

	got := SplitBranches("Dev\nMaster")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetCurrentBranch(t *testing.T) {
	t.Run("returns the current branch", func(t *testing.T) {
		want := "*Dev"

		input := []string{"", "*Dev", "Master"}
		got, _ := GetCurrentBranch(input)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("returns error as no current branch found", func(t *testing.T) {
		input := []string{"Dev", "Master"}
		_, got := GetCurrentBranch(input)

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != ErrNoCurrentBranchFound {
			t.Errorf("got '%s' want '%s'", got, ErrNoCurrentBranchFound)
		}
	})

}

func TestCheckout(t *testing.T) {

	t.Run("switched branch", func(t *testing.T) {
		want := "Switched to branch 'fake'"

		fake := fakeCommander{
			result: []byte(want),
		}

		got, _ := CheckoutBranch(fake, "fake")

		if !strings.Contains(got, want) {
			t.Errorf("wanted '%s' but got '%s'", want, got)
		}
	})

	t.Run("branch doesn't exist", func(t *testing.T) {
		want := ErrBranchDoesNotExist

		fake := fakeCommander{
			err: ErrBranchDoesNotExist,
		}

		_, got := CheckoutBranch(fake, "fake")

		if got != want {
			t.Errorf("wanted '%s' but got '%s'", want, got)
		}
	})
}
