package issue

import (
	"io/ioutil"
	"os"

	"github.com/nchern/sit/pkg/model"
)

// Create creates a new issue
func Create() error {
	if err := checkIsRepo(); err != nil {
		return err
	}
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	t := model.NewTicket()
	t.Project = cfg.ProjectName

	dir := getTicketDir(t.ID.String())
	if err := os.Mkdir(dir, defaultDirPerms); err != nil {
		return err
	}

	path := fullTicketPath(dir)
	if err := createIssueFile(path, t); err != nil {
		return err
	}

	return shellout(path).Run()
}

func createIssueFile(path string, t *model.Ticket) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.ToText(f)
}

// Edit brings an editor to edit a given issue
func Edit(partialID string) error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	found, err := findOne(partialID)
	if err != nil {
		return err
	}

	path := fullTicketPath(getTicketDir(found[0]))
	return shellout(path).Run()
}

// Delete deletes a given issue
func Delete(partialID string) error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	found, err := findOne(partialID)
	if err != nil {
		return err
	}

	return os.RemoveAll(getTicketDir(found[0]))
}

// List lists all issues to a given writer
func List() ([]*model.Ticket, error) {
	res := []*model.Ticket{}
	entries, err := ioutil.ReadDir(issuesDir)
	if err != nil {
		return res, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		id := entry.Name()
		path := fullTicketPath(getTicketDir(id))

		f, err := os.Open(path)
		if err != nil {
			return res, err
		}
		defer f.Close()

		t, err := model.ParseTicketFrom(f)
		if err != nil {
			return res, err
		}

		res = append(res, t)
	}
	return res, nil
}
