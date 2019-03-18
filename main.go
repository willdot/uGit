package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	output, _ := RunCommandWithResult("git", "branch")
	branches := splitBranches(string(output))

	for _, r := range branches {

		if r == "" {
			continue
		}
		if string(r[0]) == "*" {
			fmt.Printf("Current branch: %v\n", r)
		} else {
			fmt.Println(r)
		}
	}
}

func stdout() io.Writer {
	return os.Stdout
}

func splitBranches(s string) []string {
	result := strings.Split(s, "\n")

	return result
}
