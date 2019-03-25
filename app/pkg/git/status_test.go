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
		got, noCommit := GetFilesOrNothingToCommit(untracked)

		assertSlice(t, got, want)
		assertBool(t, noCommit, false)
	})

	t.Run("Tracked and untracked", func(t *testing.T) {
		got, noCommit := GetFilesOrNothingToCommit(trackedAndUntracked)

		assertSlice(t, got, want)
		assertBool(t, noCommit, false)
	})

	t.Run("Nothing to commit", func(t *testing.T) {
		got, noCommit := GetFilesOrNothingToCommit(nothingToCommit)

		assertSlice(t, got, nil)
		assertBool(t, noCommit, true)
	})

	t.Run("No untracked files, but changes", func(t *testing.T) {
		got, noCommit := GetFilesOrNothingToCommit(noUntrackedButChanges)

		assertSlice(t, got, nil)
		assertBool(t, noCommit, false)
	})
}

func TestGetNotStagedFiles(t *testing.T) {

	t.Run("Files not staged", func(t *testing.T) {
		want := []string{"modified:   something.go", "modified:   something/something.go"}

		got := GetNotStagedFiles(notStagedForCommit)

		assertSlice(t, got, want)
	})

	t.Run("No files un staged", func(t *testing.T) {
		want := 0

		got := GetNotStagedFiles(untracked)

		assertSliceCount(t, len(got), want)
	})

}

func assertSlice(t *testing.T, got, want []string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertBool(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func assertSliceCount(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("expected count was: %v but got: %v", want, got)
	}
}

func BenchmarkGetTrackedFiles(b *testing.B) {
	b.Run("Untracked only", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetFilesOrNothingToCommit(untracked)
		}
	})

	b.Run("Tracked and untracked", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetFilesOrNothingToCommit(trackedAndUntracked)
		}
	})

	b.Run("Nothing to commit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetFilesOrNothingToCommit(nothingToCommit)
		}
	})

	b.Run("No untracked files, but changes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetFilesOrNothingToCommit(noUntrackedButChanges)
		}
	})
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

var noUntrackedButChanges = `On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

new file:   New folder/New Text Document - Copy.txt
new file:   New folder/New Text Document.txt
new file:   a.txt
new file:   b.txt

`
var notStagedForCommit = `On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        modified:   New folder/New Text Document - Copy.txt
        modified:   one.txt
        modified:   pp/p

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

		modified:   something.go
		modified:   something/something.go

`
