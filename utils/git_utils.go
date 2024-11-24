package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetGitRoot returns the root directory of a Git repository.
func GetGitRoot(dir string) (string, error) {
	cmd := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to find Git root: %w", err)
	}
	return strings.TrimSpace(out.String()), nil
}

// GetStagedFiles returns a list of staged files in the Git repository.
func GetStagedFiles(dir string) ([]string, error) {
	cmd := exec.Command("git", "-C", dir, "diff", "--name-only", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to get staged files: %w", err)
	}
	files := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(files) == 1 && files[0] == "" {
		return []string{}, nil // No staged files
	}
	return files, nil
}
