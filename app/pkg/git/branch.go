package git

import (
	"strings"

	"uGit/app/pkg/run"

	"github.com/pkg/errors"
)

// ErrNoCurrentBranchFound is an error returned when the input branches doesn't contain a current branch indicator
var ErrNoCurrentBranchFound = errors.New("no current branch found")

// ErrBranchDoesNotExist is an error for when the branch asked to be checked out, doesn't exist
var ErrBranchDoesNotExist = errors.New("Cannot checkout branch as it doesn't exist")

// SplitBranches takes a string of branches with newline separators and splits them into a slice
func SplitBranches(s string, removeCurrent bool) []string {

	result := strings.Split(s, "\n")
	if removeCurrent {
		result = RemoveCurrentBranch(result)
	}

	return result
}

// GetCurrentBranch takes a slice of branch names and returns the current branch based on which one starts with an *
func GetCurrentBranch(branches []string) (string, error) {
	for _, r := range branches {
		if r == "" {
			continue
		}

		if string(r[0]) == "*" {
			return r, nil
		}
	}

	return "", ErrNoCurrentBranchFound
}

// GetBranches gets all local branches
func GetBranches(commander run.ICommander) (string, error) {
	result, err := run.CommandWithResult(commander)

	return result, err
}

//RemoveCurrentBranch will remove the current branch from a list of branches
func RemoveCurrentBranch(branches []string) []string {

	var result []string
	current, _ := GetCurrentBranch(branches)

	for i := 0; i < len(branches); i++ {
		if branches[i] != current && branches[i] != "" {
			result = append(result, branches[i])
		}
	}

	return result
}

// CheckoutBranch checks out a branch
func CheckoutBranch(commander run.ICommander) (string, error) {

	result, err := run.CommandWithResult(commander)

	return string(result), err
}
