package cli

import (
	"fmt"

	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func list() error {
	tickets, err := issue.List()
	if err != nil {
		return err
	}

	for _, t := range tickets {
		_, err := fmt.Printf("%s\t%s\t%s\t%s\t%s\n",
			t.State, t.ID, t.CreatedAsString(), t.User, t.Title)
		if err != nil {
			return err
		}
	}
	return nil
}
