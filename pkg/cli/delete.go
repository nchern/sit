package cli

import (
	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes issue",
	// TODO: add bulk delete
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return issue.Delete(args[0])
	},
}
