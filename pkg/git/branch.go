package git

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/willdot/uGit/pkg/run"
)

var (
	// ErrNoCurrentBranchFound is an error returned when the input branches doesn't contain a current branch indicator
	ErrNoCurrentBranchFound = errors.New("no current branch found")

	// ErrBranchDoesNotExist is an error for when the branch asked to be checked out, doesn't exist
	ErrBranchDoesNotExist = errors.New("Cannot checkout branch as it doesn't exist")
)

// GetBranches gets all local branches
func GetBranches() ([]string, error) {
	result, err := run.RunCommand("git", []string{"branch", "-a"})
	if err != nil {
		return nil, err
	}

	branches := splitBranches(result, true)

	return branches, nil
}

// RemoveRemoteOriginFromName removes the remotes/origin part of the branch
func RemoveRemoteOriginFromName(branch string) string {
	if strings.Contains(branch, "remotes/origin/") {
		x := strings.Split(branch, "remotes/origin/")

		branch = x[1]
	}

	return branch
}

func splitBranches(s string, removeCurrent bool) []string {

	branches := strings.Split(s, "\n")
	branches = trimSlice(branches)

	if removeCurrent {
		branches = removeCurrentBranch(branches)
		branches = removeOriginHead(branches)
	}

	branches = filterSlice(branches)

	return branches
}

func getCurrentBranch(branches []string) (string, error) {
	for _, branch := range branches {
		if branch == "" {
			continue
		}

		if string(branch[0]) == "*" {
			return branch, nil
		}
	}

	return "", ErrNoCurrentBranchFound
}

func removeCurrentBranch(branches []string) []string {
	var result []string
	current, _ := getCurrentBranch(branches)

	for _, branch := range branches {
		branch = RemoveRemoteOriginFromName(branch)
		if branch != current && branch != "" && branch != strings.Trim(current, "* ") {
			result = append(result, branch)
		}
	}

	return result
}

func removeOriginHead(branches []string) []string {
	var result []string

	for _, branch := range branches {
		branch = RemoveRemoteOriginFromName(branch)
		if !strings.Contains(branch, "HEAD ->") {
			result = append(result, branch)
		}
	}

	return result
}

func trimSlice(i []string) []string {
	result := make([]string, 0, len(i))

	for _, s := range i {
		result = append(result, strings.TrimSpace(s))
	}

	return result
}

func filterSlice(i []string) []string {
	uniqueValues := make(map[string]struct{})

	for _, s := range i {
		uniqueValues[s] = struct{}{}
	}

	result := make([]string, 0, len(uniqueValues))
	for k := range uniqueValues {
		result = append(result, k)
	}

	return result
}
