package issue

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Init initialises issue tracking in the current dir
func Init() error {
	return os.MkdirAll(issuesDir, defaultDirPerms)
}

// Create creates a new issue
func Create() error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	t := newTicket()
	dir := getTicketRoot(t)
	if err := os.Mkdir(dir, defaultDirPerms); err != nil {
		return err
	}

	path := fullTicketPath(dir)
	if err := createIssueFile(path, t); err != nil {
		return err
	}

	return shellout(path).Run()
}

func createIssueFile(path string, t *ticket) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.ToText(f)
}

// Edit brings an editor to edit a given issue
func Edit(partialId string) error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	found, err := findOne(partialId)
	if err != nil {
		return err
	}

	path := fullTicketPath(filepath.Join(issuesDir, found[0]))
	return shellout(path).Run()
}

// Delete deletes a given issue
func Delete(partialId string) error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	found, err := findOne(partialId)
	if err != nil {
		return err
	}

	return os.RemoveAll(filepath.Join(issuesDir, found[0]))
}

// List lists all issues to a given writer
func ListTo(w io.Writer) error {
	entries, err := ioutil.ReadDir(issuesDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		id := entry.Name()
		path := fullTicketPath(filepath.Join(issuesDir, id))

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		t, err := ParseFrom(f)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			entry.Name(), t.State, t.Created.Format(time.RFC3339), t.User); err != nil {
			return err
		}
	}
	return nil
}
