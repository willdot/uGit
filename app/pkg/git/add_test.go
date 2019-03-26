package git

import (
	"fmt"
	"testing"
	"uGit/app/pkg/run"
)

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

func ExampleAdd() {

	addCommander := run.Commander{
		Command: "git",
		Args:    []string{"add", "."},
	}

	result, err := Add(addCommander)

	fmt.Println(result)
	fmt.Println(err)

	// Output: "" ""
}
