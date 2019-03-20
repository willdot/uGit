package git

import (
	"reflect"
	"strings"
	"testing"
)

// FakeCommander is used for mocking
type FakeCommander struct {
	Result []byte
	Err    error
}

func (f *FakeCommander) RunCommand() ([]byte, error) {

	if f.Err != nil {
		return nil, f.Err
	}

	return f.Result, nil
}

func TestSplitBranch(t *testing.T) {
	t.Run("Split and keep current", func(t *testing.T) {
		want := []string{"* Current", "Dev", "Master"}

		got := SplitBranches("* Current\nDev\nMaster", false)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Split and remove current", func(t *testing.T) {
		want := []string{"Dev", "Master"}

		got := SplitBranches("* Current\n\nDev\nMaster", true)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestRemoveCurrentBranch(t *testing.T) {
	want := []string{"dev", "master"}
	input := []string{"* current", "dev", "master"}
	got := RemoveCurrentBranch(input)

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

func TestGetBranches(t *testing.T) {
	t.Run("get local branches", func(t *testing.T) {
		want := ""

		fake := &FakeCommander{
			Result: []byte(want),
		}

		got, _ := GetBranches(fake)

		if got != want {
			t.Errorf("want '%s' but got '%s'", got, want)
		}
	})
}

func TestCheckout(t *testing.T) {

	t.Run("switched branch", func(t *testing.T) {
		want := "Switched to branch 'fake'"

		fake := &FakeCommander{
			Result: []byte(want),
		}

		got, _ := CheckoutBranch(fake, "fake")

		if !strings.Contains(got, want) {
			t.Errorf("wanted '%s' but got '%s'", want, got)
		}
	})

	t.Run("branch doesn't exist", func(t *testing.T) {
		want := ErrBranchDoesNotExist

		fake := &FakeCommander{
			Err: ErrBranchDoesNotExist,
		}

		_, got := CheckoutBranch(fake, "fake")

		if got != want {
			t.Errorf("wanted '%s' but got '%s'", want, got)
		}
	})
}
