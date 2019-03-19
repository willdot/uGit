package main

import (
	"strings"

	"github.com/pkg/errors"
)

// ErrNoCurrentBranchFound is an error returned when the input branches doesn't contain a current branch indicator
var ErrNoCurrentBranchFound = errors.New("no current branch found")

// ErrBranchDoesNotExist is an error for when the branch asked to be checked out, doesn't exist
var ErrBranchDoesNotExist = errors.New("Cannot checkout branch as it doesn't exist")

// SplitBranches takes a string of branches with newline separators and splits them into a slice
func SplitBranches(s string) []string {
	result := strings.Split(s, "\n")

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

// CheckoutBranch checks out a branch
func CheckoutBranch(commander Commander, branch string) (string, error) {

	result, err := commander.combinedOutput("git", "checkout", branch)

	return string(result), err
}
