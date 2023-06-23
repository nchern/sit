package cli

import (
	"os"

	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

var issueShowCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"cat"},
	Args:    cobra.ExactArgs(1),
	Short:   "Outputs a given issue",
	RunE: func(cmd *cobra.Command, args []string) error {
		return issue.WriteByPartialID(args[0], os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(issueShowCmd)
}
