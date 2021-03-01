package model

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFrom(t *testing.T) {
	created, err := time.Parse(time.RFC3339, "2021-02-24T00:46:06+01:00")
	require.NoError(t, err)

	var tests = []struct {
		name      string
		givenPath string
		expected  *Ticket
	}{
		{"ticket",
			"./testdata/ticket.md",
			&Ticket{
				ID:      Identifier(uuid.MustParse("159693e7be3e4d2cb0e309b80098898e")),
				User:    "john",
				Created: created,
				Project: "test_prj",
				State:   ticketState("CLOSED"),
				Title:   "Expected title",
			}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.givenPath)
			require.NoError(t, err)
			defer f.Close()

			actual, err := ParseTicketFrom(f)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
