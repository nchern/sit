package cmd

import (
	"github.com/nchern/sit/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edits issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		must(issue.Edit(args[0]))
	},
}
