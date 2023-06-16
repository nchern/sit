package cli

import (
	"fmt"
	"strings"

	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output if set")
}

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list()
		},
	}

	flagVerbose bool
)

func list() error {
	tickets, err := issue.List()
	if err != nil {
		return err
	}

	for _, t := range tickets {
		fields := []string{
			string(t.State),
			t.ID.String(),
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
