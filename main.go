package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	fmt.Print("ss")

	commander := RealCommander{}
	output, err := RunCommandWithResult(commander, "git", "checkout", "master")

	if err != nil {
		fmt.Printf("error: %v", errors.WithMessage(err, ""))
		return
	}

	fmt.Println(output)
}
