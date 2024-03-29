package issue

import (
	"io"
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

// List lists all issues
func List(states ...model.TicketState) ([]*model.Ticket, error) {
	if err := checkIsRepo(); err != nil {
		return nil, err
	}

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

		t, err := loadTicket(path)
		if err != nil {
			return res, err
		}
		for _, st := range states {
			if t.State == st {
				res = append(res, t)
				break
			}
		}
	}
	return res, nil
}

// WriteByPartialID writes an issue given by partial ID to a given writer
func WriteByPartialID(partialID string, w io.Writer) error {
	if err := checkIsRepo(); err != nil {
		return err
	}

	found, err := findOne(partialID)
	if err != nil {
		return err
	}
	path := fullTicketPath(getTicketDir(found[0]))
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	return err
}

// FetchTicket fetches ticket by id
func FetchTicket(id string) (*model.Ticket, error) {
	path := fullTicketPath(getTicketDir(id))
	return loadTicket(path)
}

// Update updates a ticket
func Update(t *model.Ticket) error {
	path := fullTicketPath(getTicketDir(t.ID.String()))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.ToText(f)
}

func loadTicket(path string) (*model.Ticket, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return model.ParseTicketFrom(f)
}
