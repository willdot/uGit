package main

import "fmt"

func main() {

	fmt.Print("ss")

	output, err := RunCommandWithResult("git", "branch")

	if err != nil {
		fmt.Println(err)
		return
	}

	branches := SplitBranches(string(output))

	currentBranch, err := GetCurrentBranch(branches)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Current branch: %s", currentBranch)
}
