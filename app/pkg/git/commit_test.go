package git

import (
	"testing"
)

func TestCommit(t *testing.T) {

	want := "success"

	fake := &FakeCommander{
		Result: []byte(want),
	}

	got, _ := CommitChanges(fake)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
