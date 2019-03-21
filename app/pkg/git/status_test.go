package git

import (
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {

	want := "something"

	fake := &FakeCommander{
		Result: []byte(want),
	}

	got, _ := Status(fake)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

func TestGetTrackedFiles(t *testing.T) {
	want := []string{"something/something.go", "something/else.go", "Some folder/"}

	t.Run("Untracked only", func(t *testing.T) {
		got := GetFiles(untracked)

		assertSlice(t, got, want)
	})

	t.Run("Tracked and untracked", func(t *testing.T) {
		got := GetFiles(trackedAndUntracked)

		assertSlice(t, got, want)
	})

	t.Run("No untracked files", func(t *testing.T) {
		got := GetFiles(nothingToCommit)

		assertSlice(t, got, nil)
	})
}

func assertSlice(t *testing.T, got, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

var untracked = `On branch feature/commit
Untracked files:
(use "git add <file>..." to include in what will be committed)

something/something.go
something/else.go
Some folder/

nothing added to commit but untracked files present (use "git add" to track)`

var trackedAndUntracked = `On branch feature/commit
Your branch is ahead of 'origin/feature/commit' by 1 commit.
(use "git push" to publish your local commits)

Changes not staged for commit:
(use "git add <file>..." to update what will be committed)
(use "git checkout -- <file>..." to discard changes in working directory)

modified:   alreadyTracked.go
modified:   alsoTracked.go

Untracked files:
(use "git add <file>..." to include in what will be committed)

something/something.go
something/else.go
Some folder/

no changes added to commit (use "git add" and/or "git commit -a")`

var nothingToCommit = `On branch feature/commit
Your branch is ahead of 'origin/feature/commit' by 1 commit.
(use "git push" to publish your local commits)

nothing to commit, working tree clean`
