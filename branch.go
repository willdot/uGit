package main

import (
	"errors"
	"strings"
)

// ErrNoCurrentBranchFound is an error returned when the input branches doesn't contain a current branch indicator
var ErrNoCurrentBranchFound = errors.New("no current branch found")

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
