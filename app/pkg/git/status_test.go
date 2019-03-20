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

	got := GetUntrackedFiles(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
