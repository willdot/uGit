package git

import "testing"

func TestPush(t *testing.T) {

	want := "something"

	fake := &FakeCommander{
		Result: want,
	}

	got, _ := Push(fake)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
