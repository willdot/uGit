package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	fmt.Print("ss")

	commander := RealCommander{}
	output, err := RunCommandWithResult(commander, "git", "branch", "-a")

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		return
	}

	branches := SplitBranches(string(output))

	currentBranch, err := GetCurrentBranch(branches)

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		return
	}
	fmt.Printf("Current branch: %s", currentBranch)
}
