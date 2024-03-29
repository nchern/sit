package cli

import (
	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inits an issue repo in the current dir",
	RunE: func(cmd *cobra.Command, args []string) error {
		return issue.Init()
	},
}
