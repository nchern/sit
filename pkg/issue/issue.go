package issue

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/nchern/sit/pkg/model"
)

const (
	defaultDirPerms = 0744

	defaultEditor = "vim"

	repoRootDir      = ".sit"
	defaultIssueFile = "main.md"
	issuesDir        = repoRootDir + "/issues"
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

func getTicketRoot(t *model.Ticket) string {
	return filepath.Join(issuesDir,
		strings.ToLower(hex.EncodeToString(t.ID[:])))
}

func fullTicketPath(dir string) string {
	return filepath.Join(dir, defaultIssueFile)
}

func findOne(partialId string) ([]string, error) {
	partialId = strings.ToLower(partialId)
	entries, err := ioutil.ReadDir(issuesDir)
	if err != nil {
		return nil, err
	}

	found := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if strings.Contains(entry.Name(), partialId) {
			found = append(found, entry.Name())
		}
		if len(found) > 1 {
			return nil, fmt.Errorf("Found more than 1 issues with partial id '%s':\n%s", partialId, found)
		}
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("No issues found with ids like '%s'", partialId)
	}

	return found, nil
}
