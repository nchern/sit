package cmd

import (
	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates issue",
	Run: func(cmd *cobra.Command, args []string) {
		must(issue.Create())
	},
}
