package git

import (
	"testing"
)

func TestDeleteBranch(t *testing.T) {

	want := "success"

	fake := &FakeCommander{
		Result: []byte(want),
	}

	got, _ := DeleteBranch(fake)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
