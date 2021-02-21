package cmd

import (
	"os"

	"github.com/nchern/sit/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists issue",
	Run: func(cmd *cobra.Command, args []string) {
		must(issue.ListTo(os.Stdin))
	},
}
