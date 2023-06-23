package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nchern/sit/pkg/issue"
	"github.com/nchern/sit/pkg/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output if set")
}

var (
	listCmd = &cobra.Command{
		Aliases: []string{"ls"},
		Use:     "list",
		Short:   "lists issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list()
		},
	}

	flagVerbose bool
)

func list() error {
	const abbrevLen = 7
	tickets, err := issue.List()
	if err != nil {
		return err
	}
	sort.Sort(byDate(tickets))
	for _, t := range tickets {
		fields := []string{
			string(t.State),
			t.ID.Abbreviation(abbrevLen),
			t.Title,
		}
		if flagVerbose {
			fields = []string{
				string(t.State),
				t.ID.String(),
				t.CreatedAsString(),
				t.User,
				t.Title,
			}
		}
		if _, err := fmt.Printf("%s\n", strings.Join(fields, "\t")); err != nil {
			return err
		}
	}
	return nil
}

type byDate []*model.Ticket

func (cln byDate) Len() int { return len(cln) }

func (cln byDate) Less(i int, j int) bool { return cln[i].Created.Before(cln[j].Created) }

func (cln byDate) Swap(i int, j int) { cln[i], cln[j] = cln[j], cln[i] }
