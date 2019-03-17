package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "branch")
	//cmd.Stdout = stdout()
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	splitString(string(output))

	fmt.Printf("result:\n %v", string(output))
}

func stdout() io.Writer {
	return os.Stdout
}

func splitString(s string) {
	result := strings.Split(s, "\n")

	fmt.Println("lines")
	for _, r := range result {
		fmt.Println(r)
	}

}
