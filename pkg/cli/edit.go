package cli

import (
	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:               "edit",
	Short:             "Edits issue",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeIDs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return issue.Edit(args[0])
	},
}
