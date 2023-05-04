package git

import (
	"reflect"
	"testing"
)

// FakeCommander is used for mocking
type FakeCommander struct {
	Result       string
	ResultString []string
	Err          error
}

func (f *FakeCommander) CommandWithResult() (string, error) {

	if f.Err != nil {
		return "", f.Err
	}

	return f.Result, nil
}

func TestSplitBranch(t *testing.T) {
	t.Run("Split and keep current", func(t *testing.T) {
		want := []string{"* current", "Dev", "Master"}

		got := SplitBranches("* current\nDev\nMaster", false)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Split and remove current", func(t *testing.T) {
		want := []string{"Dev", "Master"}

		got := SplitBranches("* current\n\nDev\nMaster\nremotes/origin/current", true)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Split and remove origin head", func(t *testing.T) {
		want := []string{"Dev", "Master"}

		got := SplitBranches("* current\n\nDev\nMaster\nremotes/origin/current\nremotes/origin/HEAD -> origin/master", true)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestRemoveCurrentBranch(t *testing.T) {
	want := []string{"dev", "master"}
	input := []string{"* current", "dev", "master", "remotes/origin/current"}
	RemoveCurrentBranch(&input)

	if !reflect.DeepEqual(input, want) {
		t.Errorf("got %v want %v", input, want)
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

func TestRemoveRemoteOriginFromName(t *testing.T) {
	t.Run("Has remotes origin", func(t *testing.T) {
		input := "remotes/origin/master"

		want := "master"

		input = RemoveRemoteOriginFromName(input)

		if input != want {
			t.Errorf("want %s but got %s", want, input)
		}

	})

	t.Run("No remotes origin", func(t *testing.T) {
		input := "bug/origins"

		want := "bug/origins"

		input = RemoveRemoteOriginFromName(input)

		if input != want {
			t.Errorf("want %s but got %s", want, input)
		}

	})
}
