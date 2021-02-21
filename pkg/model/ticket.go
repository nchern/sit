package model

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ticketState string

var (
	stateOpen       ticketState = "OPEN"
	stateClosed     ticketState = "CLOSED"
	stateInProgress ticketState = "INPROGRESS"
)

/*
# Title

# Description


---
ID:         abcabcabc
State:      OPEN | CLOSED | INPROGRESS
User:       john.doe
Project:    prj
Tag:        tag;tag
*/

// Ticket represents an issue ticket in the system
//TODO: better name?
type Ticket struct {
	ID      uuid.UUID
	State   ticketState
	User    string
	Project string
	Created time.Time
	//Tag:        tag;tag

	Title       string
	Description string
}

// NewTicket creates a new ticket instance
func NewTicket() *Ticket {
	return &Ticket{
		ID:          uuid.New(),
		State:       stateOpen,
		Created:     time.Now(),
		User:        os.Getenv("USER"),
		Project:     "not implemented", // FIXME
		Title:       "<Enter title here>",
		Description: "<Enter description here>",
	}

}

// ToText writes ticket as text to a given writer
func (t *Ticket) ToText(w io.Writer) error {
	lines := []string{
		fmt.Sprintf("ID: %s", t.ID),
		fmt.Sprintf("State: %s", t.State),
		fmt.Sprintf("User: %s", t.User),
		fmt.Sprintf("Project: %s", t.Project),
		fmt.Sprintf("Created: %s", t.Created.Format(time.RFC3339)),
		"",
		"---",
		"",
		"# Title",
		t.Title,
		"",
		"# Description",
		t.Description,
	}
	for _, l := range lines {
		if _, err := fmt.Fprintln(w, l); err != nil {
			return err
		}
	}
	return nil
}

// ParseTicketFrom creates a new ticket and parses its contents from a given reader
func ParseTicketFrom(r io.Reader) (*Ticket, error) {
	t := &Ticket{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if v, found := parseField("ID: ", l); found {
			id, err := uuid.Parse(v)
			if err != nil {
				return nil, err
			}
			t.ID = id
		}
		if v, found := parseField("User: ", l); found {
			t.User = v
		}
		if v, found := parseField("State: ", l); found {
			t.State = ticketState(v)
		}
		if v, found := parseField("Created: ", l); found {
			tm, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return nil, err
			}
			t.Created = tm
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return t, nil
}

func parseField(field, s string) (string, bool) {
	if strings.HasPrefix(s, field) {
		return strings.TrimPrefix(s, field), true
	}
	return "", false
}
