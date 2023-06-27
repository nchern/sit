package issue

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	defaultDirPerms  = 0744
	defaultFilePerms = 0644

	defaultEditor = "vim"

	repoRootDir      = ".sit"
	defaultIssueFile = "main.md"
	issuesDir        = repoRootDir + "/issues"
	configFile       = repoRootDir + "/config"
)

var (
	errNotARepo = errors.New("not an issues repo: " + repoRootDir)
)

func checkIsRepo() error {
	if _, err := os.Stat(issuesDir); err != nil {
		if os.IsNotExist(err) {
			return errNotARepo
		}
		return err
	}
	return nil
}

func shellout(args ...string) *exec.Cmd {
	editorCmd := os.Getenv("EDITOR")
	if editorCmd == "" {
		editorCmd = defaultEditor
	}

	cmd := exec.Command(editorCmd, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func getTicketDir(id string) string {
	return filepath.Join(issuesDir, id)
}

func fullTicketPath(dir string) string {
	return filepath.Join(dir, defaultIssueFile)
}

func findOne(partialID string) ([]string, error) {
	partialID = strings.ToLower(partialID)
	entries, err := ioutil.ReadDir(issuesDir)
	if err != nil {
		return nil, err
	}

	found := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if strings.Contains(entry.Name(), partialID) {
			found = append(found, entry.Name())
		}
		if len(found) > 1 {
			return nil, fmt.Errorf("Found more than 1 issues with partial id '%s':\n%s", partialID, found)
		}
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("No issues found with ids like '%s'", partialID)
	}

	return found, nil
}

// FindIDs looks up ids matching provides substring
func FindIDs(partialID string) ([]string, error) {
	partialID = strings.ToLower(partialID)
	entries, err := ioutil.ReadDir(issuesDir)
	if err != nil {
		return nil, err
	}

	found := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if strings.Contains(entry.Name(), partialID) {
			found = append(found, entry.Name())
		}
	}

	return found, nil
}
