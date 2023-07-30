package cli

import (
	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates issue",
	RunE: func(cmd *cobra.Command, args []string) error {
		return issue.Create()
	},
}
