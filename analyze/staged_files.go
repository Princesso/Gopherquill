package analyze

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetStagedFiles retrieves staged files in the current Git directory.
func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get staged files: %w", err)
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	return files, nil
}
