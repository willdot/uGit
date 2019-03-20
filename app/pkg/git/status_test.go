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

func TestGetUntrackedFiles(t *testing.T) {

	want := []string{"something/something.go", "something/else.go", "Some folder/"}

	input := `On branch feature/commit
Untracked files:
(use "git add <file>..." to include in what will be committed)

something/something.go
something/else.go
Some folder/

nothing added to commit but untracked files present (use "git add" to track)`

	got := GetFiles(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetUntrackedFilesNoUntrackedFiles(t *testing.T) {

	input := `On branch feature/commit
Your branch is ahead of 'origin/feature/commit' by 1 commit.
(use "git push" to publish your local commits)

nothing to commit, working tree clean`

	got := GetFiles(input)

	if got != nil {
		t.Errorf("got %v want %v", got, nil)
	}
}

func TestGetTrackedAndUntrackedFiles(t *testing.T) {

	want := []string{"something/something.go", "something/else.go", "Some folder/"}

	input := `On branch feature/commit
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

	got := GetFiles(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
