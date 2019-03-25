package git

import "testing"

func TestAdd(t *testing.T) {

	want := "something"

	fake := &FakeCommander{
		Result: want,
	}

	got, _ := Add(fake)

	if got != want {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
